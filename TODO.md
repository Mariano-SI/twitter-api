# TODO

Roadmap de features pendentes. Marque com `[x]` conforme for fechando.

---

## 1. Imagens em posts (em andamento)

Infra já pronta: bucket R2, `storage.Storage` interface, adapter `r2storage`, wiring no `main.go`, tabela `post_images`.

Falta:
- [ ] Repository de `post_images` (`Create`, `ListByPostID`, `DeleteByPostID`)
- [x] Alterar `CreatePostDto` para aceitar imagens (`multipart/form-data` no handler, não JSON)
- [x] No handler `Create`, ler arquivos do `multipart.Form` e validar:
  - [x] Máx. N imagens por post (sugestão: 4, igual ao Twitter)
  - [x] Mime type permitido (`image/jpeg`, `image/png`, `image/webp`)
  - [x] Tamanho máximo por arquivo (sugestão: 5MB)
- [x] No `postService.Create`:
  - [x] Gerar `key` única por imagem (`posts/{post_id}/{uuid}.{ext}`)
  - [x] Upload pro R2 via `imageStorage.Upload`
  - [x] Insert na tabela `post_images` com `url` retornada e `position`
  - [x] **Tudo dentro de uma transaction** — se o insert falhar depois do upload, fazer cleanup das keys já enviadas pro R2 (ou aceitar órfãos e ter um job de limpeza)
- [ ] Decidir estratégia de rollback: o R2 não participa da transaction SQL — pensar em compensação manual

---

## 2. CRUD de posts (faltam read/update/delete)

- [ ] `GET /posts/:id` — retorna post + imagens associadas (JOIN com `post_images`)
- [ ] `GET /posts` — lista paginada (cursor ou offset?). Considerar filtros: `?user_id=`
- [ ] `PATCH /posts/:id` — editar `content` (só o autor pode; checar `user_id == ctx.userId`)
- [ ] `DELETE /posts/:id` — soft delete (preenche `deleted_at`)
  - [ ] Decidir: ao soft-deletar post, apagar as imagens do R2? (sugestão: manter por 30 dias e ter job)
- [ ] Validar `content` no DTO: não vazio + máx. 280 chars

---

## 3. Feed

- [ ] Nova migration: tabela `follows` (ver seção 4)
- [ ] `GET /feed` — posts dos usuários que o user logado segue, ordenados por `created_at DESC`, paginado
- [ ] Decidir: o que retornar se o user não segue ninguém? (sugestão: posts mais recentes globais como discovery)
- [ ] Performance: index em `posts(user_id, created_at)` e em `follows(follower_id)`

---

## 4. Follows

- [ ] Migration: tabela `follows`
  ```sql
  follower_id  CHAR(36) NOT NULL  -- quem segue
  following_id CHAR(36) NOT NULL  -- quem é seguido
  created_at   TIMESTAMP
  PRIMARY KEY (follower_id, following_id)
  FK ambos -> users(id) ON DELETE CASCADE
  CHECK (follower_id <> following_id)  -- não pode seguir a si mesmo
  ```
- [ ] `POST /users/:id/follow` — seguir
- [ ] `DELETE /users/:id/follow` — deixar de seguir
- [ ] `GET /users/:id/followers` — paginado
- [ ] `GET /users/:id/following` — paginado
- [ ] Bloquear auto-follow no service também (defesa em profundidade)

---

## 5. Likes em posts

Tabela `post_likes` já existe.

- [ ] Repository de `post_likes`
- [ ] `POST /posts/:id/like` — idempotente (não erro se já curtiu)
- [ ] `DELETE /posts/:id/like` — idempotente
- [ ] Incluir contagem de likes no payload de `GET /posts/:id` e do feed

---

## 6. Comments

Tabelas `comments` e `comment_likes` já existem.

- [ ] CRUD básico: criar, listar por post, deletar
- [ ] Like em comment (mesma lógica de post like)

---

## Backlog / ideias

- [ ] Paginação por cursor (não offset) — escala melhor
- [ ] Rate limiting (ex: máx. 10 posts/min por user)
- [ ] Job de limpeza de imagens órfãs no R2 (posts deletados ou uploads falhos)
- [ ] Health check endpoint (`GET /health`)
- [ ] Logging estruturado (substituir `log.Println` por algo tipo `slog`)
- [ ] Testes (unit no service, integration no repository com DB real)
- [ ] Migrar dev de `r2.dev` (rate-limited) pra custom domain quando for pra produção

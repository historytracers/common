# Agent Instructions

- All new source files must have `// SPDX-License-Identifier: GPL-3.0-or-later` as the first line.
- When modifying a Go datatype file (`src/go/data-type.go`) or a Java datatype file (`src/android/.../com/historytracers/common/*.java`), the counterpart must always be updated in sync.
- All files must use LF (`\n`) line endings, never CRLF (`\r\n`). Enforced by `.gitattributes` (`* text=auto eol=lf`).

## Creating Smartphone Screen Files

Smartphone screen files are stored at `src/smartphone/??-??/uuid.json`, where `??-??` is the language code (e.g. `en-US`, `pt-BR`, `es-ES`) and `uuid.json` is shared across all languages.

### Source Material

- The **Original Text** for each lesson is at `historytracers/lang/??-??/<uuid>.json`.
- Each Original Text contains sections identified by `id` fields (e.g. `SECTION_prerequisites`, `SECTION_introduction`, etc.). Use only the sections explicitly requested; do not include content from sections marked as excluded.
- The Original Text also lists academic sources in its `source` arrays. These UUIDs must be carried over into the smartphone `source_menu`.

### File Structure

Each smartphone JSON file follows this schema (match existing files exactly):

```
{
   "title": "...",
   "index": "...",
   "sources": ["<original-text-uuid>"],
   "license": ["SPDX-License-Identifier: GPL-3.0-or-later", "CC BY-NC 4.0 DEED"],
   "last_update": ["<unix-timestamp>"],
   "authors": "",
   "reviewers": "",
   "version": 1,
   "type": "sm_game",
   "content": [ ... screens ... ]
}
```

Each screen object in `content`:

```
{
   "id": "<uuid>",
   "audio": "",
   "smile": "<emoji-name>",
   "source_menu": [ ... ],
   "text": [ ... ],
   "desc": null,
   "next": "<next-screen-uuid>",
   "prev": "",
   "jumpTo": "",
   "answer": null | "yes" | "no",
   "score": 1 | 2,
   "position": ""
}
```

### Screen Creation Rules

1. **Smiley emojis**: Each screen must have a `"smile"` field. Follow the progression pattern used by existing files: `nerd` (intro) -> `inlove` / `happy` (explanations) -> `shocking` (surprising content) -> `thinking` (questions) -> `party` (conclusion).
2. **Questions**: Include at least one Yes/No question screen per lesson. Set `"answer"` to `"yes"` or `"no"`, and `"score"` to `2`. Question screens have empty `"source_menu": []`.
3. **Navigation**: The `"next"` field of each screen points to the `"id"` of the following screen. The last screen has `"next": ""`.
4. **Multiple text entries**: A screen may have more than one entry in `"text"` (e.g. a markdown description followed by an HTML image or SVG). Each entry has `"format": "markdown"` or `"format": "html"`.

### Source Menu Rules

- Every non-question screen should have a `"source_menu"` array referencing the academic sources from the Original Text.
- Each source entry: `{"type": 0|1, "uuid": "<source-uuid>", "text": "<citation>", "page": "<url>", "date_time": {"type": "gregory", "year": "...", "month": "", "day": ""}}`.
- **The `"page"` field must always be a URL.** Look up the `src_url` from the `sources` table in `historytracers/lang/sources/history_tracers.db` using the source UUID as `src_id`. Never use page numbers, volume labels, or other non-URL text in this field.
- Always include an "Original Text" entry pointing to the lesson UUID with its historytracers.org URL.
- **Translate `"text"` labels** for each language (e.g. "Texto Original" for pt-BR and es-ES, not "Original Text").

### Encoding

- Files must be valid JSON with **actual UTF-8 characters** (e.g. `ñ`, `é`, `ã`, `¿`), not Unicode escape sequences (`\u00XX`).
- When writing files programmatically, use `json.dump(data, f, ensure_ascii=False)` to preserve real characters.
- The `write` tool may convert Unicode escapes to literal text; prefer using a Python script with `json.dump` for files containing non-ASCII characters.
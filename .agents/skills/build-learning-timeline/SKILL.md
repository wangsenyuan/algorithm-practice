---
name: build-learning-timeline
description: Use when creating, refreshing, or appending a daily learning timeline for newly added competitive-programming solution packages in learn-go.
---

# Build Learning Timeline

Record only packages selected by the repository's deterministic discovery tool, then advance its Git boundary only after the record is complete.

## Workflow

1. Set `REPO_ROOT` to the learn-go repository root. Run exactly:

   ```bash
   python3 "$REPO_ROOT/tools/timeline/discover.py" --repo "$REPO_ROOT"
   ```

   Treat any Git/state/discovery error as a hard stop and write nothing. Package selection comes only from `packages`; never substitute Git status, mtimes, or a different commit range.

2. For every emitted package, read every path in its `files` array. Read additional artifacts only when those files directly reference them. Never edit anything in a solution package.

3. From that evidence, draft one compact, natural-prose entry per package, linking its directory as `../../src/.../`. Capture the most useful learning point without repeating the statement. Algorithm, proof, complexity spike, rejected flaw, and pitfall are optional observations, not mandatory labels. Never invent criticism. If evidence is insufficient, report the package as unrecorded, make no timeline/index edits, and do not advance the marker.

4. Use the local date for `docs/timeline/YYYY-MM-DD.md` and the exact heading:

   ```markdown
   # Learning Timeline — YYYY-MM-DD
   ```

   Preserve existing content and append only missing package links. Never rewrite an earlier date. If discovery emits no packages, do not create a dated file.

5. Only after all summaries are complete, update `docs/timeline/README.md`. It must start with `# Learning Timeline`, contain exactly:

   ```markdown
   <!-- through-commit: FULL_HEAD -->
   ```

   where `FULL_HEAD` is discovery's 40-character `head`, and list dated-file links in descending date order. With no packages, advance a valid marker without creating a dated file.

6. Rerun the exact discovery command. Success requires `packages` to be `[]`. If not, repair omissions or malformed links rather than duplicating entries. Never run `indexgen`.

## Report

Report changed files, recorded and skipped/unrecorded packages, and the final discovery verification. On any failure, state it plainly and preserve the last valid marker.

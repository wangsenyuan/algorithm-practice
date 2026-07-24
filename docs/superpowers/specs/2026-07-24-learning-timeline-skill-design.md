# Learning Timeline Skill Design

Date: 2026-07-24

## Goal

Add a repository-local Codex skill that turns newly added solution packages into
short, durable learning notes. Each invocation creates or updates only today's
timeline file. The generated solution index links to the timeline but does not
aggregate or rewrite its summaries.

## Scope

The feature consists of:

- a repository-local `build-learning-timeline` skill;
- a Cursor `/build-learning-timeline` command;
- a shared deterministic discovery script under `tools/timeline`;
- dated notes under `docs/timeline/YYYY-MM-DD.md`;
- a lightweight `docs/timeline/README.md`;
- one timeline reference in `docs/index/README.md`, rendered by `tools/indexgen`.

It does not run automatically, summarize modified packages, or require every
entry to contain a fixed set of learning labels.

## Skill location and trigger

Store the skill at:

```text
.agents/skills/build-learning-timeline/
```

Its description should trigger when the user asks to build, refresh, or append
the daily learning timeline from newly added contest-solution packages.

The skill owns the judgment-heavy part of the workflow: reading source material
and writing concise, evidence-based summaries. Mechanical Git discovery stays
in `tools/timeline/discover.py` so repeated invocations and editor integrations
select packages consistently.

## Cursor command

Store the Cursor version at:

```text
.cursor/commands/build-learning-timeline.md
```

It is an explicit custom command, invoked as `/build-learning-timeline`, because
timeline generation happens only when requested. It is not an always-applied or
file-attached Cursor rule.

The command follows the same discovery, summary, output, state, and failure
workflow as the Codex skill. It calls `tools/timeline/discover.py` rather than
carrying a second implementation.

## Discovery boundary

The discovery script returns newly added solution packages from three sources:

1. Files added by commits after the last recorded `through-commit`, through
   current `HEAD`.
2. Files currently added to the Git index.
3. Untracked, non-ignored files.

Only directories whose newly added files include `solution.go` qualify as new
solution packages. Modified, deleted, renamed-only, generated-index, and
unrelated files are excluded.

On the first run, when no `through-commit` exists, find the most recent commit
reachable from `HEAD` that added a `solution.go`, and use that commit's parent
as the baseline. Consequently, an unrelated documentation or tooling commit at
`HEAD` cannot hide the latest committed solution package. The first run includes
that solution-adding commit, any later committed additions, and current staged
and untracked additions.

The script emits deterministic JSON containing:

- the baseline and current `HEAD`;
- candidate package paths;
- relevant source-document paths within each package;
- whether each package came from committed, staged, or untracked additions.

Packages are sorted by repository-relative path.

## Deduplication and state

`docs/timeline/README.md` contains a machine-readable HTML comment:

```markdown
<!-- through-commit: FULL_COMMIT_SHA -->
```

All dated timeline files are scanned for package links. A package already linked
from a timeline file is excluded from future discovery. This prevents an
untracked package recorded today from appearing again after it is committed.

Advance `through-commit` to current `HEAD` only after the timeline update
succeeds. Advancing it is still useful when a commit contains no qualifying new
package, because later runs should not rescan that commit.

## Summary workflow

For every discovered package, the skill reads:

- `solution.go`;
- `solution_test.go`, when present;
- `problem.md`, `readme.md`, or `README.md`, when present;
- other directly relevant package artifacts only when the primary files refer
  to them.

It writes one compact entry per package. The entry should capture the most
useful learning point supported by the files, which may include:

- the central algorithm or technique;
- a complexity or proof breakthrough;
- an implementation spike or difficult edge case;
- a flaw in an earlier idea or implementation;
- a reusable pitfall.

These are prompts for judgment, not mandatory fields. The skill must omit
dimensions unsupported by the source and must not invent a flaw merely to
complete a template. “Flaw” normally refers to a rejected draft, performance
problem, correctness gap, or test-discovered issue—not an unsupported criticism
of the final solution.

Use natural prose rather than labels such as `Technique:` or `Flaw:` unless a
label materially improves clarity.

## Dated timeline format

Each invocation creates or updates:

```text
docs/timeline/YYYY-MM-DD.md
```

The date is the local invocation date, not a commit date. If the previous
timeline run was several days ago, all newly discovered packages still go into
today's single file.

Example:

```markdown
# Learning Timeline — 2026-07-24

- [Codeforces 513D2](../../src/codeforces/set0/set5/set510/set513/d2/)
  — Uses preorder intervals and a minimum-ending greedy construction. The key
  spike was replacing D1's cubic interval DP with an explicit-stack `O(n+c)`
  DFS; recursion depth and formatted output were important D2 pitfalls.
```

When today's file already exists, preserve its content and append only missing
package entries. Do not rewrite earlier dated files.

## Timeline index

`docs/timeline/README.md` contains:

- the `through-commit` state comment;
- links to dated files in descending date order.

It contains no copied package summaries. The dated files remain the source of
the learning content.

`tools/indexgen` adds one stable link from `docs/index/README.md`:

```markdown
[Learning timeline](../timeline/README.md)
```

`indexgen` does not parse, aggregate, regenerate, or delete timeline content.

## Failure behavior

- If there are no new packages, do not create an empty dated file. Advance the
  valid commit marker and report that nothing needed summarizing.
- If Git discovery fails or the stored commit is invalid, stop with a clear
  diagnostic and do not update timeline state.
- If a package lacks readable evidence for a useful summary, report it and leave
  it unrecorded rather than writing filler.
- If writing the dated file or timeline index fails, do not advance the commit
  marker.
- Preserve existing manual text in today's file and all earlier timeline files.

## Testing

### Discovery script

Use temporary Git repositories to verify:

- first-run discovery from the parent of the latest solution-adding commit;
- first-run behavior when later commits contain no solution additions;
- later `through-commit..HEAD` discovery;
- staged and untracked additions;
- exclusion of modifications, deletions, ignored files, and unrelated files;
- grouping by package and stable ordering;
- deduplication using existing timeline links;
- invalid-state failure without mutation.

### Cursor command

Forward-test the command against a clean temporary repository containing
committed, staged, untracked, and modified packages. It must select the same
packages and produce the same timeline structure as the Codex skill. Add a
static check that the command remains user-invoked and references the shared
discovery script.

### Skill behavior

Before writing the skill, run a baseline scenario without it and record where an
agent's package selection or summary format is inconsistent. After writing the
skill, repeat the same scenario and verify:

- correct package selection from the discovery output;
- one concise entry per package;
- evidence-based optional learning dimensions;
- preservation and idempotency of existing timeline files;
- state advancement only after successful output.

Validate the finished skill metadata with the standard skill validator.

### Index generator

Add a focused test proving that the generated index home links to the timeline.
Run the existing `tools/indexgen` tests to ensure platform indexes remain
unchanged and deterministic.

## First-run result for the current repository

With current `HEAD` at `71ddbbc06`, the most recent solution-adding commit is
`8cc81d778066866d909bb12f04fe216b26bb4bea`. The first run would therefore
discover Codeforces 2241F from that commit and ten untracked packages:

- AtCoder ABC426 C, D, E, F, and G;
- Codeforces 212D, 493D, 513D1, 513D2, and 742B.

It would create `docs/timeline/2026-07-24.md`, create the timeline index, record
the full current `HEAD`, and leave every solution package itself unchanged.

# CF2236G First Example SVG Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Create a self-contained SVG that shows the first sample test case's tree and all three query paths without exposing query answers or solution ideas.

**Architecture:** The SVG uses reusable definitions for the tree's neutral edges and labeled nodes. Three side-by-side panels draw a query-specific directional path between those reusable layers, keeping each query visually independent.

**Tech Stack:** SVG 1.1-compatible XML, CSS media query for light/dark appearance, `xmllint`, and `rsvg-convert`.

---

### Task 1: Create and verify the example diagram

**Files:**
- Create: `src/codeforces/set2/set22/set223/set2236/g/first-example.svg`
- Reference: `src/codeforces/set2/set22/set223/set2236/g/problem.md`

- [ ] **Step 1: Create the SVG**

Create `src/codeforces/set2/set22/set223/set2236/g/first-example.svg` with this complete content:

```svg
<svg xmlns="http://www.w3.org/2000/svg" width="1200" height="500" viewBox="0 0 1200 500" role="img" aria-labelledby="title desc">
  <title id="title">First sample tree and its three query paths</title>
  <desc id="desc">Three panels repeat a four-node tree. Each panel highlights one directed query path: 1 to 4, 2 to 3, and 2 to 4. Node labels include friendliness values.</desc>

  <defs>
    <style>
      :root {
        color-scheme: light dark;
        --bg: #ffffff;
        --panel: #f8fafc;
        --ink: #172033;
        --muted: #5f6b7a;
        --edge: #b8c1cc;
        --border: #d9e0e8;
        --path: #2563eb;
        --path-soft: #dbeafe;
      }
      @media (prefers-color-scheme: dark) {
        :root {
          --bg: #111827;
          --panel: #182235;
          --ink: #f3f6fb;
          --muted: #b7c0cd;
          --edge: #637083;
          --border: #344258;
          --path: #60a5fa;
          --path-soft: #1e3a5f;
        }
      }
      .canvas { fill: var(--bg); }
      .panel { fill: var(--panel); stroke: var(--border); stroke-width: 2; }
      .panel-title { fill: var(--ink); font: 600 22px system-ui, sans-serif; text-anchor: middle; }
      .panel-subtitle { fill: var(--muted); font: 15px system-ui, sans-serif; text-anchor: middle; }
      .edge { fill: none; stroke: var(--edge); stroke-linecap: round; stroke-width: 5; }
      .active-halo { fill: none; stroke: var(--path-soft); stroke-linecap: round; stroke-linejoin: round; stroke-width: 15; }
      .active-path { fill: none; marker-end: url(#arrow); stroke: var(--path); stroke-linecap: round; stroke-linejoin: round; stroke-width: 7; }
      .node { fill: var(--bg); stroke: var(--ink); stroke-width: 2.5; }
      .node-id { fill: var(--ink); font: 600 18px system-ui, sans-serif; text-anchor: middle; }
      .node-value { fill: var(--muted); font: 13px system-ui, sans-serif; text-anchor: middle; }
      .legend { fill: var(--muted); font: 15px system-ui, sans-serif; }
      .legend-edge { stroke: var(--edge); stroke-linecap: round; stroke-width: 5; }
      .legend-path { stroke: var(--path); stroke-linecap: round; stroke-width: 7; }
    </style>

    <marker id="arrow" viewBox="0 0 12 12" refX="10" refY="6" markerWidth="9" markerHeight="9" orient="auto-start-reverse">
      <path d="M 1 1 L 11 6 L 1 11 z" fill="var(--path)"/>
    </marker>

    <g id="tree-edges">
      <path class="edge" d="M 180 150 L 75 300"/>
      <path class="edge" d="M 180 150 L 180 300"/>
      <path class="edge" d="M 180 150 L 285 300"/>
    </g>

    <g id="tree-nodes">
      <g transform="translate(180 150)">
        <circle class="node" r="31"/>
        <text class="node-id" y="-3">1</text>
        <text class="node-value" y="17">a=0</text>
      </g>
      <g transform="translate(75 300)">
        <circle class="node" r="31"/>
        <text class="node-id" y="-3">2</text>
        <text class="node-value" y="17">a=0</text>
      </g>
      <g transform="translate(180 300)">
        <circle class="node" r="31"/>
        <text class="node-id" y="-3">3</text>
        <text class="node-value" y="17">a=4</text>
      </g>
      <g transform="translate(285 300)">
        <circle class="node" r="31"/>
        <text class="node-id" y="-3">4</text>
        <text class="node-value" y="17">a=1</text>
      </g>
    </g>
  </defs>

  <rect class="canvas" width="1200" height="500"/>

  <g transform="translate(20 20)">
    <rect class="panel" width="373" height="390" rx="18"/>
    <text class="panel-title" x="186.5" y="43">Query 1</text>
    <text class="panel-subtitle" x="186.5" y="68">1 → 4</text>
    <use href="#tree-edges"/>
    <path class="active-halo" d="M 199 177 L 267 275"/>
    <path class="active-path" d="M 199 177 L 267 275"/>
    <use href="#tree-nodes"/>
  </g>

  <g transform="translate(413 20)">
    <rect class="panel" width="373" height="390" rx="18"/>
    <text class="panel-title" x="186.5" y="43">Query 2</text>
    <text class="panel-subtitle" x="186.5" y="68">2 → 3</text>
    <use href="#tree-edges"/>
    <path class="active-halo" d="M 93 275 L 162 177 L 176 266"/>
    <path class="active-path" d="M 93 275 L 162 177 L 176 266"/>
    <use href="#tree-nodes"/>
  </g>

  <g transform="translate(806 20)">
    <rect class="panel" width="373" height="390" rx="18"/>
    <text class="panel-title" x="186.5" y="43">Query 3</text>
    <text class="panel-subtitle" x="186.5" y="68">2 → 4</text>
    <use href="#tree-edges"/>
    <path class="active-halo" d="M 93 275 L 180 150 L 267 275"/>
    <path class="active-path" d="M 93 275 L 180 150 L 267 275"/>
    <use href="#tree-nodes"/>
  </g>

  <g transform="translate(386 455)">
    <line class="legend-path" x1="0" y1="0" x2="48" y2="0"/>
    <text class="legend" x="62" y="5">query path</text>
    <line class="legend-edge" x1="238" y1="0" x2="286" y2="0"/>
    <text class="legend" x="300" y="5">other tree edge</text>
  </g>
</svg>
```

- [ ] **Step 2: Validate the XML**

Run:

```bash
xmllint --noout src/codeforces/set2/set22/set223/set2236/g/first-example.svg
```

Expected: exit status `0` with no output.

- [ ] **Step 3: Render a PNG preview**

Run:

```bash
rsvg-convert \
  src/codeforces/set2/set22/set223/set2236/g/first-example.svg \
  -o /tmp/cf2236g-first-example.png
```

Expected: exit status `0` and a non-empty `/tmp/cf2236g-first-example.png`.

- [ ] **Step 4: Inspect the rendered preview**

Open `/tmp/cf2236g-first-example.png` with the local image viewer and verify:

- all three panels fit without clipping;
- node IDs and values match the first sample;
- each highlighted route matches its panel's query direction;
- arrowheads do not cover node labels;
- no query answer or algorithmic explanation appears.

- [ ] **Step 5: Check whitespace and repository scope**

Run:

```bash
git diff --no-index --check /dev/null \
  src/codeforces/set2/set22/set223/set2236/g/first-example.svg
git status --short
```

Expected: no whitespace errors. The status output lists the new SVG within the already-untracked problem package; implementation must not stage or modify unrelated files.

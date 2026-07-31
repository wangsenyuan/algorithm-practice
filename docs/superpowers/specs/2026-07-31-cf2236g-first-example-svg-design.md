# CF2236G First Example SVG Design

## Goal

Create a standalone SVG that illustrates the first test case from
`src/codeforces/set2/set22/set223/set2236/g/problem.md` without revealing query
answers, solution ideas, or algorithmic hints.

## Content

- Show three side-by-side panels, one for each query:
  - `Query 1: 1 → 4`
  - `Query 2: 2 → 3`
  - `Query 3: 2 → 4`
- Repeat the same four-node tree in every panel:
  - edges: `1–2`, `1–3`, `1–4`
  - friendliness values: `a₁=0`, `a₂=0`, `a₃=4`, `a₄=1`
- Label every node with both its vertex number and friendliness value.
- Highlight only the path belonging to the panel's query.
- Keep non-path edges visible in a subdued neutral style.

## Presentation

- Use a clean, readable layout that remains legible when embedded in Markdown.
- Use arrowheads or endpoint emphasis to preserve the query direction.
- Pair color with line weight so path highlighting remains understandable
  without relying on color alone.
- Include a compact legend for highlighted and non-path edges.
- Do not display hospitable-subsegment counts or any derivation of them.

## Deliverable

Save the final asset as:

`src/codeforces/set2/set22/set223/set2236/g/first-example.svg`

Validate that the SVG is well-formed XML and visually inspect a rendered copy
before completion.

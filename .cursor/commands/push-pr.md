# push-pr

Ask the user: "What branch name do you want to use?"

Then follow these steps:

1. Run `git status` and `git diff --staged` to understand what changes exist.

2. Run `git log -5 --oneline` to understand the recent commit style.

3. **If the branch name is `main`:**
   - Stage relevant changed files and commit directly on main with an appropriate commit message.
   - Push: `git push origin main`
   - Done — no PR needed.

4. **If the branch name is anything other than `main`:**
   - Create and checkout the branch: `git checkout -b <branch>` (or `git checkout <branch>` if it already exists).
   - Stage relevant changed files and commit with an appropriate commit message.
   - Push and set upstream: `git push -u origin <branch>`
   - Analyze all commits and changes on the branch vs main, then create a PR:
     ```
     gh pr create --title "<short title under 70 chars>" --body "$(cat <<'EOF'
     ## Summary
     <1-3 bullet points describing what changed and why>

     ## Test plan
     <bulleted checklist of what to verify>
     EOF
     )"
     ```
   - Merge the PR and delete the remote branch:
     ```
     gh pr merge --merge --delete-branch
     ```
   - Report the PR URL to the user.

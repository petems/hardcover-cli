# Markdown Documentation Website Validation with GitHub Actions

## Overview

This document outlines a comprehensive validation strategy for Markdown-based documentation websites using GitHub Actions. The goal is to ensure high quality, consistency, and accuracy across all documentation files through automated checks that run on every pull request and push to the main branch.

## Validation Categories

### 1. Spelling and Grammar

#### Why It's Important
- **Professional credibility**: Spelling and grammar errors undermine the perceived quality and reliability of documentation
- **User experience**: Clear, error-free content improves comprehension and reduces confusion
- **Maintainability**: Consistent language standards make content easier to update and maintain
- **Accessibility**: Proper grammar and spelling improve compatibility with screen readers and translation tools

#### Implementation in GitHub Actions

**Primary Tool: `cspell` (Code Spell Checker)**
```yaml
- name: Check Spelling
  uses: streetsidesoftware/cspell-action@v6
  with:
    files: "docs/**/*.md"
    config: ".cspell.json"
    incremental_files_only: false
```

**Alternative: `pyspelling`**
```yaml
- name: Check Spelling with PySpelling
  run: |
    pip install pyspelling
    pyspelling -c .pyspelling.yml
```

**Grammar Checking with `vale`**
```yaml
- name: Vale Linting
  uses: errata-ai/vale-action@reviewdog
  with:
    files: docs
    vale_flags: "--config=.vale.ini"
  env:
    GITHUB_TOKEN: ${{secrets.GITHUB_TOKEN}}
```

#### Configuration Considerations

**cspell Configuration (`.cspell.json`)**:
```json
{
  "version": "0.2",
  "language": "en",
  "words": [
    "GitHub",
    "API",
    "JSON",
    "TypeScript"
  ],
  "dictionaries": [
    "technical-terms",
    "project-specific"
  ],
  "ignorePaths": [
    "node_modules/**",
    "**/*.min.js"
  ],
  "ignoreRegExpList": [
    "\\b[A-Z]{2,}\\b",
    "```[\\s\\S]*?```"
  ]
}
```

**Vale Configuration (`.vale.ini`)**:
```ini
StylesPath = .vale/styles
MinAlertLevel = suggestion
Packages = Microsoft, write-good

[*.md]
BasedOnStyles = Vale, Microsoft, write-good
```

#### Reporting Failures
- **PR Comments**: Vale and cspell can post detailed comments on specific lines
- **Workflow Summary**: Generate a summary report with total error counts
- **Annotations**: Use GitHub's annotation system to highlight issues in the Files Changed tab

### 2. Link Validation (404 Checks)

#### Why It's Important
- **User experience**: Broken links frustrate users and interrupt their workflow
- **SEO impact**: Search engines penalize sites with broken links
- **Content maintenance**: Dead links indicate outdated content that needs review
- **Professional image**: Functional links demonstrate attention to detail and quality

#### Implementation in GitHub Actions

**Primary Tool: `markdown-link-check`**
```yaml
- name: Check Markdown Links
  uses: gaurav-nelson/github-action-markdown-link-check@v1
  with:
    use-quiet-mode: 'yes'
    use-verbose-mode: 'yes'
    config-file: '.markdown-link-check.json'
    folder-path: 'docs/'
```

**Alternative: `lychee` (Rust-based, faster)**
```yaml
- name: Link Checker
  uses: lycheeverse/lychee-action@v1.10.0
  with:
    args: --verbose --no-progress 'docs/**/*.md'
    fail: true
  env:
    GITHUB_TOKEN: ${{secrets.GITHUB_TOKEN}}
```

**Custom Script for Internal Links**
```yaml
- name: Check Internal Links
  run: |
    find docs -name "*.md" -exec grep -l "\[.*\](.*\.md)" {} \; | \
    xargs python scripts/check_internal_links.py
```

#### Configuration Considerations

**markdown-link-check Configuration (`.markdown-link-check.json`)**:
```json
{
  "ignorePatterns": [
    {
      "pattern": "^http://localhost"
    },
    {
      "pattern": "^https://example.com"
    }
  ],
  "httpHeaders": [
    {
      "urls": ["https://api.github.com/"],
      "headers": {
        "Accept": "application/vnd.github.v3+json",
        "User-Agent": "markdown-link-check"
      }
    }
  ],
  "timeout": "20s",
  "retryOn429": true,
  "retryCount": 3,
  "fallbackHttpHeaders": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:96.0) Gecko/20100101 Firefox/96.0"
  }
}
```

**Rate Limiting Considerations**:
- Implement delays between external link checks
- Use caching for frequently accessed domains
- Set appropriate timeouts for slow-responding sites
- Consider running external link checks on a schedule rather than every PR

#### Reporting Failures
- **Detailed logs**: Show which links failed and why (404, timeout, etc.)
- **File-by-file breakdown**: Group errors by source file
- **Retry suggestions**: Indicate if failures might be temporary

### 3. Markdown Linting and Formatting

#### Why It's Important
- **Consistency**: Uniform formatting improves readability and professional appearance
- **Maintainability**: Consistent structure makes bulk updates easier
- **Tool compatibility**: Proper Markdown syntax ensures compatibility with various renderers
- **Collaboration**: Style standards reduce merge conflicts and review overhead

#### Implementation in GitHub Actions

**Primary Tool: `markdownlint`**
```yaml
- name: Markdown Linting
  uses: DavidAnson/markdownlint-cli2-action@v16
  with:
    globs: 'docs/**/*.md'
    config: '.markdownlint.json'
```

**Alternative: `remark-lint`**
```yaml
- name: Remark Lint
  run: |
    npm install -g remark-cli remark-lint
    remark docs/ --use remark-lint --frail
```

**Formatting with Prettier**
```yaml
- name: Check Markdown Formatting
  run: |
    npx prettier --check "docs/**/*.md"
```

#### Configuration Considerations

**markdownlint Configuration (`.markdownlint.json`)**:
```json
{
  "default": true,
  "MD013": {
    "line_length": 100,
    "code_blocks": false,
    "tables": false
  },
  "MD033": {
    "allowed_elements": ["br", "details", "summary"]
  },
  "MD041": false,
  "MD046": {
    "style": "fenced"
  }
}
```

**Key Rules to Consider**:
- **MD013**: Line length limits (with exceptions for code and tables)
- **MD025**: Single H1 per document
- **MD033**: HTML usage restrictions
- **MD040**: Code block language specification
- **MD046**: Code block style consistency (fenced vs. indented)

#### Reporting Failures
- **Line-specific annotations**: Point to exact locations of violations
- **Rule explanations**: Include links to rule documentation
- **Auto-fix suggestions**: Where possible, suggest specific corrections

### 4. General Content Quality & Best Practices

#### Why It's Important
- **Accessibility**: Proper alt text and semantic structure improve screen reader compatibility
- **SEO**: Well-structured content with appropriate headings improves search visibility
- **Usability**: Consistent terminology and clear navigation enhance user experience
- **Maintenance**: Quality standards prevent technical debt accumulation

#### Implementation in GitHub Actions

**Accessibility Checks**
```yaml
- name: Check Accessibility
  run: |
    python scripts/accessibility_check.py docs/
```

**Content Structure Validation**
```yaml
- name: Validate Document Structure
  run: |
    # Check for proper heading hierarchy
    find docs -name "*.md" -exec python scripts/heading_structure.py {} \;
    
    # Verify all images have alt text
    grep -r "!\[.*\]" docs/ | grep -v "!\[.*\](" && exit 1 || true
```

**Terminology Consistency**
```yaml
- name: Check Terminology
  uses: errata-ai/vale-action@reviewdog
  with:
    files: docs
    vale_flags: "--config=.vale.ini --ext=.md"
```

#### Specific Quality Checks

**1. Image Alt Text Validation**
```python
# scripts/check_alt_text.py
import re
import sys
import glob

def check_alt_text(file_path):
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Find images without alt text
    images_without_alt = re.findall(r'!\[\s*\]\([^)]+\)', content)
    
    if images_without_alt:
        print(f"Missing alt text in {file_path}: {images_without_alt}")
        return False
    return True
```

**2. Heading Structure Validation**
```python
# scripts/heading_structure.py
import re

def validate_heading_structure(content):
    headings = re.findall(r'^(#{1,6})\s+(.+)$', content, re.MULTILINE)
    
    if not headings:
        return True
    
    # Check for single H1
    h1_count = sum(1 for h, _ in headings if len(h) == 1)
    if h1_count != 1:
        return False
    
    # Check for proper hierarchy (no skipping levels)
    levels = [len(h) for h, _ in headings]
    for i in range(1, len(levels)):
        if levels[i] > levels[i-1] + 1:
            return False
    
    return True
```

**3. Content Length and Readability**
```yaml
- name: Content Quality Checks
  run: |
    # Check for overly long paragraphs
    python scripts/readability_check.py docs/
    
    # Verify code blocks have language specified
    grep -r "^```$" docs/ && exit 1 || true
```

#### Configuration Considerations

**Custom Terminology Dictionary**:
```yaml
# .vale/styles/Terminology/accept.txt
API
GitHub
JavaScript
TypeScript
Markdown
```

**Content Guidelines**:
- Maximum paragraph length (e.g., 4-5 sentences)
- Required sections for certain document types
- Consistent code example formatting
- Standard file naming conventions

#### Reporting Failures

**Comprehensive Reporting Strategy**:
```yaml
- name: Generate Quality Report
  if: always()
  run: |
    echo "## Documentation Quality Report" >> $GITHUB_STEP_SUMMARY
    echo "### Spelling Errors: $SPELL_ERRORS" >> $GITHUB_STEP_SUMMARY
    echo "### Broken Links: $LINK_ERRORS" >> $GITHUB_STEP_SUMMARY
    echo "### Linting Issues: $LINT_ERRORS" >> $GITHUB_STEP_SUMMARY
    echo "### Accessibility Issues: $A11Y_ERRORS" >> $GITHUB_STEP_SUMMARY
```

## Complete GitHub Actions Workflow Example

```yaml
name: Documentation Validation

on:
  pull_request:
    paths:
      - 'docs/**'
  push:
    branches:
      - main
    paths:
      - 'docs/**'

jobs:
  validate-docs:
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout
      uses: actions/checkout@v4
      with:
        fetch-depth: 0

    - name: Setup Node.js
      uses: actions/setup-node@v4
      with:
        node-version: '18'

    - name: Setup Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.11'

    # Parallel validation jobs
    - name: Spell Check
      uses: streetsidesoftware/cspell-action@v6
      with:
        files: "docs/**/*.md"
        config: ".cspell.json"

    - name: Markdown Lint
      uses: DavidAnson/markdownlint-cli2-action@v16
      with:
        globs: 'docs/**/*.md'
        config: '.markdownlint.json'

    - name: Link Check
      uses: lycheeverse/lychee-action@v1.10.0
      with:
        args: --verbose --no-progress 'docs/**/*.md'
        fail: true

    - name: Vale Lint
      uses: errata-ai/vale-action@reviewdog
      with:
        files: docs
        vale_flags: "--config=.vale.ini"
      env:
        GITHUB_TOKEN: ${{secrets.GITHUB_TOKEN}}

    - name: Custom Quality Checks
      run: |
        python scripts/accessibility_check.py docs/
        python scripts/content_quality.py docs/

    - name: Generate Report
      if: always()
      run: |
        echo "## Documentation Validation Summary" >> $GITHUB_STEP_SUMMARY
        echo "All validation checks completed. See individual step results above." >> $GITHUB_STEP_SUMMARY
```

## Best Practices and Additional Considerations

### Performance Optimization
- **Caching**: Cache spell check dictionaries and Vale styles
- **Incremental checks**: Only validate changed files in PRs when possible
- **Parallel execution**: Run independent checks simultaneously
- **Conditional execution**: Skip expensive checks for draft PRs

### Integration with Development Workflow
- **Pre-commit hooks**: Run basic checks locally before pushing
- **VS Code extensions**: Integrate tools into the editor for real-time feedback
- **Scheduled runs**: Perform comprehensive external link checks weekly
- **Branch protection**: Require validation passes before merging

### Customization and Maintenance
- **Regular updates**: Keep action versions and tool configurations current
- **Team training**: Ensure contributors understand validation requirements
- **Exception handling**: Provide clear processes for legitimate rule exceptions
- **Feedback loop**: Regularly review and adjust rules based on team feedback

This comprehensive validation strategy ensures high-quality documentation while maintaining developer productivity through automation and clear feedback mechanisms.
## Description

Please provide a clear and concise description of the changes in this PR.

### Type of Change

Please delete options that are not relevant:

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Code refactoring (no functional changes)
- [ ] Test improvements
- [ ] CI/CD improvements

## Related Issues

Fixes #(issue_number)
Closes #(issue_number)
Related to #(issue_number)

## Changes Made

Please describe the changes made in this PR:

- 
- 
- 

## Testing

### Test Coverage

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] All existing tests pass
- [ ] New tests pass

### Manual Testing

Please describe the manual testing performed:

- [ ] Tested on Linux
- [ ] Tested on macOS
- [ ] Tested on Windows
- [ ] Tested CLI commands
- [ ] Tested error scenarios

### Test Commands

```bash
# Commands used for testing
go test ./...
go test -race ./...
./hardcover --help
# Add specific test commands here
```

## Code Quality

- [ ] Code follows Go best practices
- [ ] Code is properly formatted (`gofmt`)
- [ ] Code passes linting (`golangci-lint`)
- [ ] Code includes proper error handling
- [ ] Code includes appropriate comments/documentation

## Security

- [ ] No sensitive information is exposed
- [ ] Input validation is implemented where needed
- [ ] Security best practices are followed
- [ ] Dependencies are up to date and secure

## Performance

- [ ] Changes do not negatively impact performance
- [ ] Benchmarks run (if applicable)
- [ ] Memory usage is reasonable

## Documentation

- [ ] Code is self-documenting
- [ ] README updated (if needed)
- [ ] API documentation updated (if needed)
- [ ] Help text updated (if needed)
- [ ] CHANGELOG updated (if needed)

## Dependencies

- [ ] No new dependencies added
- [ ] New dependencies are necessary and well-maintained
- [ ] Dependencies are properly versioned
- [ ] `go.mod` and `go.sum` are updated

## Breaking Changes

If this PR introduces breaking changes, please describe them and the migration path:

- 
- 

## Screenshots/Examples

If applicable, add screenshots or example outputs to help explain your changes:

```
# Example command output
$ hardcover --help
...
```

## Checklist

- [ ] I have read the [CONTRIBUTING](CONTRIBUTING.md) guidelines
- [ ] My code follows the project's coding standards
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published

## Additional Notes

Add any additional notes or context about the PR here.

## Reviewer Focus Areas

Please pay special attention to:

- [ ] Logic correctness
- [ ] Error handling
- [ ] Performance implications
- [ ] Security considerations
- [ ] API design
- [ ] Test coverage

---

**For Reviewers:**

- [ ] Code review completed
- [ ] Tests verified
- [ ] Documentation reviewed
- [ ] Security considerations checked
- [ ] Performance impact assessed
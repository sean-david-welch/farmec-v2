# Bugs & Pending Work

## 1. Form Validation & Error Handling — Warranty Claim & Machine Registration

Public-facing support forms currently have no client-side validation or server-side error feedback.

**What's needed:**
- Server-side: add `clean()` / `clean_<field>()` methods to `WarrantyclaimForm` and `MachineregistrationForm`, return field-level errors
- Template: render `{{ form.field.errors }}` (or equivalent) next to each input so the user knows what to fix
- Required field indicators and user-friendly error messages
- Consider basic JS validation (HTML5 `required`, `pattern`) as a first line of defence

---

## 2. SEO — Identify Target Search Queries

Need to define the keyword set the site should rank for before any on-page or meta work.

**What's needed:**
- Compile a list of target queries (e.g. "farm machinery Ireland", supplier/machine-specific terms, local dealer queries)
- Map queries to specific pages (supplier detail, machine detail, blog posts, etc.)
- See `SEO.md` for any existing notes — expand it once the query list is agreed

---

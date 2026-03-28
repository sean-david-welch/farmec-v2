# Feature Requests

## 1. Warranty Claim Image Uploads
Allow customers to attach photos when submitting a warranty claim (e.g. damaged parts, serial number plates).

**Implementation notes:**
- Add an `ImageField` (or `FileField`) to `Warrantyclaim` model, stored on S3
- Support multiple images per claim — likely a separate `WarrantyclaimImage` model with a FK to `Warrantyclaim`
- Render a file input in the public warranty claim form dialog
- Display attached images in the admin detail view (Unfold inline)
- Consider file type and size validation on both client and server side

---

## 2. YouTube Video Embeds on Blog Posts
Allow staff to attach a YouTube video to a blog post that renders as an embedded player.

**Implementation notes:**
- Add a `youtube_url` `URLField` (nullable) to the `Blog` model
- Extract the video ID from the URL in a model property or template tag for use in the embed iframe
- Render the player conditionally in `blog_detail.html` when the field is populated
- Validate that the URL is a recognised YouTube format (`youtube.com/watch`, `youtu.be`) in the model's `clean()` method

---

## 3. Staff Expense Tracker with CSV Export
An admin-only tool for staff to log expenses with receipt attachments, exportable to CSV for accounting.

**Implementation notes:**
- Create a new `expenses/` Django app (or add to `support/`) with an `Expense` model:
  - `staff` (FK to `User`)
  - `date` (`DateField`)
  - `description` (`CharField`)
  - `amount` (`DecimalField`)
  - `category` (`CharField` with choices, e.g. fuel, accommodation, equipment, other)
  - `receipt` (`FileField` → S3)
  - Inherits `BaseModel` for `uid`, `created`, `modified`
- Unfold admin with list view, filters by staff/date/category, and receipt preview
- CSV export action on the admin changelist (`ExportActionMixin` or a custom `ModelAdmin.action`)
- Exported columns: date, staff name, category, description, amount, receipt URL
- Restrict access to staff/admin users only — no public-facing view needed

# SEO Notes

## Ranking Context
- "farmec ireland" — ranking #1 ✓
- "farmec" — unwinnable; Romanian cosmetics company (est. 1889, Wikipedia page) dominates
- Target keywords: supplier brand + ireland (e.g. "SIP mower ireland", "MX loader ireland", "Twose ireland"), "farm machinery ireland", "agricultural machinery importer ireland"

---

## Supplier Backlink Status
Checked May 2026. Priority: email those with existing distributor pages first.

| Supplier | Links to Farmec? | Notes |
|---|---|---|
| BFS | ✅ Yes | `bfs.uk.com/stockists/farmec-ireland-ltd/` |
| SIP Slovenia | ✅ Yes | Listed on dealer map at `sip.si` |
| Woodbay Turf | ✅ Yes | `woodbayturftech.com/distributor/farmec-ireland-ltd` — links to `farmecireland.com` not `farmec.ie`, needs updating |
| Ovlac | ❌ No | Has a distributors page — email to be added |
| MX-Mailleux | ❌ No | Has a dealer finder map — email to be added |
| TeeJet | ❌ No | Has a distributor locator — email to be added |
| Falc | ❌ No | Distributes to 60+ countries, no Farmec mention |
| Aerway / Salford | ❌ No | No distributor page visible |
| Annovi Reverberi | ❌ No | No distributor page visible |
| Arag GPS | ❌ No | No distributor page visible |
| Bargam | ❌ No | No distributor page visible |
| Rogers Sprayer | ❌ No | No distributor page visible |
| Twose | ❌ No | No distributor page visible |

---

## Pending Improvements

### 1. Machine SEO copy (High Priority)
`meta_title` and `meta_description` fields added to `Machine` model but not yet populated. Same approach as suppliers — query DB, search each machine, generate and apply copy.

### 2. Backlinks from supplier websites (High Priority — no code needed)
Email Ovlac, MX-Mailleux, TeeJet, Falc to add Farmec to their distributor/dealer pages. These have existing infrastructure.
Also email Woodbay Turf to update their link from `farmecireland.com` → `farmec.ie`.

### 3. Google Business Profile (High Priority — no code needed)
Ensure Farmec is listed, fully filled out with correct categories ("Agricultural Machinery Dealer"), photos, and reviews enabled. Local SEO is a separate ranking layer.

### 4. Internal Linking — Supplier → Spare Parts (Medium Priority)
Supplier detail pages don't link to their corresponding spare parts page. Hurts both SEO (crawlability) and UX.

**Implementation:**
- Add a "Spare Parts" link on `catalog/supplier_detail.html` pointing to the supplier's spare parts page

### 5. Blog Content (Ongoing)
Each blog post is an opportunity to rank for specific brand/product searches.

**Target topics:**
- "{Supplier name} machinery ireland"
- "{Machine type} for sale ireland"
- "farm machinery warranty ireland"
- "agricultural spare parts ireland"

### 6. Supplier Descriptions (Content Review)
Descriptions should mention Ireland, Farmec, and what the supplier is known for. Thin or manufacturer-copied descriptions hurt ranking.

### 7. FAQ Schema on supplier pages (Low Priority)
Add FAQ JSON-LD to supplier pages with questions like "Where can I buy SIP machinery in Ireland?" — can appear as rich results directly in Google.

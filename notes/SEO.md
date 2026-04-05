# SEO Notes

## Current State
- Meta titles: unique per page ✓
- Meta descriptions: improved for supplier list, spare parts list/detail ✓
- Supplier detail: uses `supplier.description` dynamically ✓
- OG/Twitter tags: present on all key pages ✓
- Google Analytics: configured (G-SENSREH8YD) ✓
- Canonical URLs: present in base.html ✓

## Ranking Context
- "farmec ireland" — ranking #1 ✓
- "farmec" — unwinnable; Romanian cosmetics company (est. 1889, Wikipedia page) dominates
- Target keywords: supplier brand + ireland (e.g. "SIP mower ireland", "MX loader ireland", "Twose ireland"), "farm machinery ireland", "agricultural machinery importer ireland"

---

## Pending Improvements

### 1. robots.txt + sitemap.xml (High Priority)
Currently getting 404s from bots on robots.txt. Without a sitemap, Google has to discover pages by crawling links alone.

**Implementation:**
- Add `django.contrib.sitemaps` to `INSTALLED_APPS`
- Create `farmec/sitemaps.py` with sitemap classes for `Supplier`, `Machine`, `Blog`, `Exhibition`, `SpareParts`
- Add sitemap and robots.txt URLs to `farmec/urls.py`
- robots.txt should point to sitemap URL

### 2. Structured Data / JSON-LD (High Priority)
Tells Google explicitly what the business is. Can enable rich results in search.

**Implementation:**
- Add `LocalBusiness` JSON-LD to homepage (`pages/home.html`):
  - name, url, telephone, address (Clonross, Drumree, Co. Meath), geo coordinates
- Add `Product` JSON-LD to machine detail pages (`catalog/machine_detail.html`)
- Add `ItemList` JSON-LD to supplier list page

### 3. Machine Detail Meta Descriptions (Medium Priority)
Machine detail pages have unique titles but no description override — falling back to the generic base.html description.

**Implementation:**
- Add `{% block meta %}` to `catalog/machine_detail.html` using `machine.description` + "Farmec Ireland" + "Co. Meath"

### 4. Internal Linking — Supplier → Spare Parts (Medium Priority)
Supplier detail pages don't link to their corresponding spare parts page. This hurts both SEO (crawlability) and UX.

**Implementation:**
- Add a "Spare Parts" button/link on `catalog/supplier_detail.html` pointing to the supplier's spare parts page

### 5. Blog Content (Ongoing)
Each blog post is an opportunity to rank for specific brand/product searches.

**Target topics:**
- "{Supplier name} machinery ireland"
- "{Machine type} for sale ireland"
- "farm machinery warranty ireland"
- "agricultural spare parts ireland"

### 6. Supplier Descriptions (Content Review)
If descriptions in the DB are thin or copied from manufacturer sites, Google may not rank supplier pages well. Each description should mention Ireland, Farmec, and what the supplier is known for.

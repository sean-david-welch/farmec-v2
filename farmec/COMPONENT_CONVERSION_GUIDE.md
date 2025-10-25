# React to Phoenix Component Conversion Guide

This document outlines the conversion of React components to Phoenix HEEx templates and LiveView components.

## ✅ Completed Conversions

### Simple Presentational Components
All converted components are located in:
- `lib/farmec_web/components/custom_components.ex`
- `lib/farmec_web/components/display_info.ex`
- `assets/js/hooks.js`

#### 1. SocialLinks Component
**Location:** `lib/farmec_web/components/custom_components.ex`
**Usage:**
```elixir
<.social_links
  facebook="https://facebook.com/farmec"
  twitter="https://twitter.com/farmec"
  instagram="https://instagram.com/farmec"
  linkedin="https://linkedin.com/farmec"
  website="https://farmec.ie"
  youtube="https://youtube.com/farmec"
/>
```

#### 2. ToTopButton Component
**Location:** `lib/farmec_web/components/custom_components.ex`
**Usage:**
```elixir
<.to_top_button />
```
**Note:** Requires the `ScrollToTop` hook in `assets/js/hooks.js` (already configured)

#### 3. TimelineCard Component
**Location:** `lib/farmec_web/components/custom_components.ex`
**Usage:**
```elixir
<.timeline_card timeline={timeline} is_admin={@is_admin}>
  <!-- Admin actions go here as slot content -->
  <.delete_button id={timeline.id} resource_key="timelines" />
</.timeline_card>
```

#### 4. DisplaysInfo Data
**Location:** `lib/farmec_web/components/display_info.ex`
**Usage:**
```elixir
# Get data
specials = FarmecWeb.DisplayInfo.specials_items()
stats = FarmecWeb.DisplayInfo.stats_items()

# Render with helper component
<%= for item <- FarmecWeb.DisplayInfo.specials_items() do %>
  <FarmecWeb.DisplayInfo.info_card item={item} />
<% end %>
```

#### 5. Map Component
**Location:** `lib/farmec_web/components/custom_components.ex`
**Usage:**
```elixir
<.google_map
  lat={53.49200990196934}
  lng={-6.5423895598058435}
  width="600px"
  height="600px"
/>
```
**Requirements:**
- Google Maps API key must be loaded in layout
- Add to `root.html.heex`:
```html
<script src="https://maps.googleapis.com/maps/api/js?key=<%= Application.get_env(:farmec, :google_maps_key) %>"></script>
```

#### 6. DownloadPdf Component
**Location:** `lib/farmec_web/components/custom_components.ex`
**Usage:**
```elixir
<.download_pdf_button type="warranty" id={warranty.id} />
<.download_pdf_button type="registration" id={registration.id} />
```
**Server Requirements:**
- Create route: `get "/api/pdf/:type/:id", PdfController, :download`
- Implement PDF generation in controller

---

## ⚠️ Components Requiring LiveView Implementation

These components need server-side state management and should be implemented as LiveView components or LiveView event handlers.

### 1. DeleteButton Component
**Original:** `client/src/components/DeleteButton.tsx`

**Current Status:** Basic component created in `custom_components.ex`, but needs LiveView integration.

**Implementation Strategy:**
```elixir
# In your LiveView module (e.g., TimelineLive.ex)

def handle_event("delete_resource", %{"id" => id, "resource" => resource}, socket) do
  case YourContext.delete_resource(resource, id) do
    {:ok, _} ->
      {:noreply,
       socket
       |> put_flash(:info, "#{resource} deleted successfully")
       |> push_navigate(to: ~p"/")}

    {:error, _changeset} ->
      {:noreply, put_flash(socket, :error, "Failed to delete #{resource}")}
  end
end
```

**Usage in templates:**
```elixir
<.delete_button
  id={item.id}
  resource_key="timelines"
  phx-click="delete_resource"
  phx-target={@myself}
/>
```

### 2. AccountButton Component
**Original:** `client/src/components/AccountButton.tsx`

**Complexity:** High - requires authentication state and conditional rendering

**Implementation Strategy:**
This should be implemented as part of your navigation layout. In Phoenix, you typically handle authentication in the router and pass `@current_user` to templates.

**Example Implementation:**
```elixir
# In your layout component or LiveView
defmodule FarmecWeb.Components.Navigation do
  use Phoenix.Component
  use Phoenix.VerifiedRoutes, endpoint: FarmecWeb.Endpoint

  attr :current_user, :map, default: nil

  def account_dropdown(assigns) do
    ~H"""
    <li class="nav-item">
      <%= if @current_user do %>
        <p class="nav-list-item">Account</p>
      <% else %>
        <a href={~p"/login"} class="nav-list-item">Login</a>
      <% end %>

      <ul class="nav-drop">
        <%= if @current_user do %>
          <%= if @current_user.role == :admin do %>
            <li class="nav-drop-item">
              <a href={~p"/users"}>Users</a>
            </li>
            <li class="nav-drop-item">
              <a href={~p"/carousels"}>Carousels</a>
            </li>
            <li class="nav-drop-item">
              <a href={~p"/line-items"}>Line Items</a>
            </li>
          <% end %>
          <li class="nav-drop-item">
            <a href={~p"/warranty"}>Warranty Claims</a>
          </li>
          <li class="nav-drop-item">
            <a href={~p"/registrations"}>Registrations</a>
          </li>
          <li class="nav-drop-item">
            <button phx-click="logout">Logout</button>
          </li>
        <% end %>
      </ul>
    </li>
    """
  end
end
```

### 3. MobileLogin Component
**Original:** `client/src/components/MobileLogin.tsx`

**Complexity:** High - similar to AccountButton but for mobile layout

**Implementation Strategy:**
Similar to AccountButton, but adapted for mobile sidebar navigation. Can reuse the same authentication logic.

```elixir
# In your sidebar component
defmodule FarmecWeb.Components.Sidebar do
  use Phoenix.Component
  use Phoenix.VerifiedRoutes, endpoint: FarmecWeb.Endpoint

  attr :current_user, :map, default: nil
  attr :on_click, JS, default: %JS{}

  def mobile_account_nav(assigns) do
    ~H"""
    <%= if @current_user do %>
      <%= if @current_user.role == :admin do %>
        <a class="nav-item" href={~p"/users"} phx-click={@on_click}>
          <i class="fa-solid fa-users"></i>
          Users
        </a>
        <!-- More admin links -->
      <% end %>
      <a class="nav-item" href={~p"/warranty"} phx-click={@on_click}>
        <i class="fa-solid fa-pen-to-square"></i>
        Warranty Claims
      </a>
      <button class="nav-item" phx-click="logout">
        Logout
      </button>
    <% else %>
      <a href={~p"/login"} class="nav-item" phx-click={@on_click}>
        <i class="fa-solid fa-user-circle"></i>
        Login
      </a>
    <% end %>
    """
  end
end
```

---

## 🔧 Required Setup Steps

### 1. Font Awesome Icons
The React app used `@fortawesome/react-fontawesome`. For Phoenix, you should include Font Awesome via CDN or npm.

**Add to `root.html.heex`:**
```html
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" />
```

**Or install via npm:**
```bash
cd assets
npm install --save @fortawesome/fontawesome-free
```

Then import in `app.css`:
```css
@import "@fortawesome/fontawesome-free/css/all.css";
```

### 2. Google Maps API Key
Store in `config/config.exs`:
```elixir
config :farmec,
  google_maps_key: System.get_env("GOOGLE_MAPS_API_KEY") || "your-api-key"
```

### 3. Authentication Setup
If not already configured, you'll need:
- User authentication context (consider using `mix phx.gen.auth`)
- Current user assignment in LiveView socket/connection
- Role-based authorization helpers

### 4. CSS Migration
You'll need to migrate the CSS modules from React to Phoenix:
- `Utils.module.css` → Phoenix CSS or Tailwind utilities
- `Suppliers.module.css` → Component-specific CSS
- `About.module.css` → Page-specific CSS
- `Home.module.css` → Page-specific CSS
- `Header.module.css` → Layout CSS
- `Sidebar.module.css` → Layout CSS

Consider using Phoenix's built-in CSS organization:
- `assets/css/app.css` - Global styles
- Component-specific classes directly in HEEx templates
- Or continue using CSS modules with appropriate build tooling

---

## 📝 Next Steps

1. **Create LiveView modules** for pages that need interactivity
2. **Implement authentication** system if not already present
3. **Migrate CSS** from React modules to Phoenix asset pipeline
4. **Create controller actions** for PDF downloads and other API endpoints
5. **Test each component** in the Phoenix context
6. **Set up routes** for all converted components

---

## 🎯 File Structure Reference

```
farmec/
├── lib/
│   └── farmec_web/
│       ├── components/
│       │   ├── core_components.ex        # Phoenix default components
│       │   ├── custom_components.ex      # Converted React components
│       │   ├── display_info.ex           # Static data module
│       │   └── layouts.ex                # Layout components
│       ├── controllers/
│       │   └── pdf_controller.ex         # Create this for PDF downloads
│       └── live/
│           └── (your_live_views)         # LiveView modules for interactive pages
└── assets/
    └── js/
        ├── app.js                         # Main JS entry (updated with hooks)
        └── hooks.js                       # Phoenix LiveView hooks (created)
```

---

## 💡 Tips for Using Components

1. **Import components in LiveView:**
```elixir
import FarmecWeb.CustomComponents
import FarmecWeb.DisplayInfo
```

2. **Use in templates:**
```heex
<.social_links facebook="..." twitter="..." />
<.to_top_button />
<.google_map lat={53.492} lng={-6.542} />
```

3. **For authenticated content:**
```elixir
<%= if @current_user do %>
  <.delete_button id={item.id} resource_key="items" />
<% end %>
```

4. **Server-side vs Client-side:**
- Use LiveView for state management and user interactions
- Use hooks for pure client-side behavior (scroll, maps initialization)
- Use regular links/forms for simple navigation and submissions

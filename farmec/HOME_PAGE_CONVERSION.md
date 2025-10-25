# Home Page Conversion - React to Phoenix

This document explains the conversion of the React Home page to Phoenix server-side rendering.

## ✅ What Was Converted

### Original React Files:
- `client/src/pages/Home.tsx`
- `client/src/templates/Carousel.tsx`
- `client/src/templates/Displays.tsx`
- `client/src/templates/Contact.tsx`

### New Phoenix Files:

1. **Controller**: `lib/farmec_web/controllers/page_controller.ex`
   - Fetches mock carousel images
   - Gets static data from `DisplayInfo` module
   - Renders home template with data

2. **Components**: `lib/farmec_web/components/custom_components.ex`
   - `carousel/1` - Image carousel with navigation
   - `displays_section/1` - Stats and specials display
   - `contact_section/1` - Contact information and map

3. **Template**: `lib/farmec_web/controllers/page_html/home.html.heex`
   - Full page HTML with SEO meta tags
   - Uses all three main components
   - Includes Font Awesome CDN
   - Server-side rendered

4. **JavaScript**: `assets/js/hooks.js`
   - `Carousel` hook - Handles slide navigation and auto-play
   - `ScrollToTop` hook - Smooth scroll functionality

## 🎯 Key Differences from React

### React (Client-Side):
```jsx
const Home = () => {
  const { data: carousels, isLoading, isError } = useGetResource('carousels');

  if (isError) return <ErrorPage />;
  if (isLoading) return <Loading />;

  return (
    <>
      <Carousel images={images}/>
      <Displays/>
      <Contact/>
    </>
  );
};
```

### Phoenix (Server-Side):
```elixir
def home(conn, _params) do
  carousel_images = [...]  # From DB or static
  specials = FarmecWeb.DisplayInfo.specials_items()
  stats = FarmecWeb.DisplayInfo.stats_items()

  render(conn, :home,
    carousel_images: carousel_images,
    specials: specials,
    stats: stats
  )
end
```

## 📋 How to Use

### 1. Start the Phoenix Server:
```bash
cd /Users/seanwelch/farmec-v2/farmec
mix phx.server
```

### 2. Visit the Home Page:
```
http://localhost:4000/
```

### 3. What You'll See:
- **Hero Carousel** with navigation buttons and auto-play
- **Stats Section** ("Farmec At A Glance")
- **Specials Section** ("What Can We Offer")
- **Contact Section** with map and business info
- **Scroll-to-Top Button**

## 🔧 Customization

### Change Carousel Images:
Edit `lib/farmec_web/controllers/page_controller.ex`:
```elixir
carousel_images = [
  "/images/your-image-1.jpg",
  "/images/your-image-2.jpg",
  "/images/your-image-3.jpg"
]
```

Or fetch from database:
```elixir
carousel_images = Farmec.Content.list_carousels()
  |> Enum.map(& &1.image)
```

### Change Stats/Specials:
Edit `lib/farmec_web/components/display_info.ex`:
```elixir
def specials_items do
  [
    %{
      title: "Your Feature",
      description: "Description here",
      icon: "fa-solid fa-icon-name",
      link: "/your-link"
    }
  ]
end
```

### Carousel Auto-Play Speed:
Edit `assets/js/hooks.js`:
```javascript
startAutoPlay() {
  this.autoPlayInterval = setInterval(() => {
    this.nextSlide();
  }, 5000);  // Change 5000 to desired milliseconds
}
```

### Disable Auto-Play:
Comment out this line in `assets/js/hooks.js`:
```javascript
// this.startAutoPlay();
```

## 🎨 Styling

Currently using placeholder CSS classes. You need to add your styles:

### Option 1: Add to `assets/css/app.css`
```css
.hero-container { /* styles */ }
.slideshow { /* styles */ }
.slides { /* styles */ }
.fade-in { opacity: 1; transition: opacity 1s; }
.fade-out { opacity: 0; transition: opacity 1s; }
/* etc... */
```

### Option 2: Copy from React CSS Modules
Copy styles from your React app:
- `client/src/styles/Carousel.module.css`
- `client/src/styles/Info.module.css`
- `client/src/styles/Home.module.css`
- `client/src/styles/Utils.module.css`

Convert `.className` to just `className` in your CSS.

## 📦 Component API

### Carousel Component:
```heex
<.carousel images={list_of_image_urls} class="optional-class" />
```

### Displays Section:
```heex
<.displays_section
  stats={list_of_stat_items}
  specials={list_of_special_items}
  class="optional-class"
/>
```

### Contact Section:
```heex
<.contact_section class="optional-class">
  <!-- Form or other content goes in slot -->
  <div>Contact form here</div>
</.contact_section>
```

## 🚀 Next Steps

1. **Add Real Database Data**:
   - Create `Carousels` schema and context
   - Replace mock data in controller
   - Query from database

2. **Add Contact Form**:
   - Create `ContactForm` component
   - Handle form submission
   - Send emails or save to DB

3. **Add CSS Styling**:
   - Port CSS from React modules
   - Or create new styles from scratch

4. **Add Navigation/Header**:
   - Create header component with nav menu
   - Add to layout or template

5. **Add Footer**:
   - Create footer component
   - Include in layout

6. **Add Error/Loading States**:
   - Handle when images fail to load
   - Show loading indicator if needed

## 💡 Benefits of This Approach

### Server-Side Rendering:
- ✅ **Better SEO** - Search engines see full HTML
- ✅ **Faster Initial Load** - No JavaScript bundle to download
- ✅ **Progressive Enhancement** - Works without JS
- ✅ **Simpler Mental Model** - No client state management

### When Pages Reload:
- Full page refresh on navigation
- No complex routing logic
- Standard browser behavior
- Perfect for mostly-static content

### Easy to Add Interactivity:
If you need real-time features later:
- Convert specific components to LiveView
- Keep server-rendering for static parts
- Best of both worlds

## 🔍 Troubleshooting

### Images Not Loading:
Place images in `priv/static/images/` directory

### Carousel Not Working:
Check browser console for JS errors. Ensure hooks are loaded in `app.js`.

### Components Not Found:
Make sure `import FarmecWeb.CustomComponents` is in `page_html.ex`

### Styles Not Applied:
Add CSS to `assets/css/app.css` and restart server with `mix phx.server`

## 📚 Related Files

- **Router**: `lib/farmec_web/router.ex` (defines the `/` route)
- **HTML Components**: `lib/farmec_web/components/custom_components.ex`
- **Static Data**: `lib/farmec_web/components/display_info.ex`
- **JavaScript Hooks**: `assets/js/hooks.js`
- **App JS Entry**: `assets/js/app.js`

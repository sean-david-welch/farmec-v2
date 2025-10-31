defmodule FarmecWeb.PageLive do
  use FarmecWeb, :live_view
  import FarmecWeb.CustomComponents

  @impl true
  def mount(_params, _session, socket) do
    # Mock carousel images (replace with DB query later)
    carousel_images = [
      "/images/default.jpg",
      "/images/default.jpg",
      "/images/default.jpg"
    ]

    # Get display info from our static module
    specials = FarmecWeb.DisplayInfo.specials_items()
    stats = FarmecWeb.DisplayInfo.stats_items()

    socket =
      socket
      |> assign(:carousel_images, carousel_images)
      |> assign(:specials, specials)
      |> assign(:stats, stats)
      |> assign(:page_title, "Home - Farmec Ireland")

    {:ok, socket}
  end
end

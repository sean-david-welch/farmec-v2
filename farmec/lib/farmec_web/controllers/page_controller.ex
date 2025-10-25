defmodule FarmecWeb.PageController do
  use FarmecWeb, :controller

  def home(conn, _params) do
    # Mock carousel images (replace with DB query later)
    carousel_images = [
      "/images/carousel1.jpg",
      "/images/carousel2.jpg",
      "/images/carousel3.jpg"
    ]

    # Get display info from our static module
    specials = FarmecWeb.DisplayInfo.specials_items()
    stats = FarmecWeb.DisplayInfo.stats_items()

    render(conn, :home,
      layout: false,
      carousel_images: carousel_images,
      specials: specials,
      stats: stats,
      page_title: "Home - Farmec Ireland"
    )
  end
end

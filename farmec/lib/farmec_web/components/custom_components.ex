defmodule FarmecWeb.CustomComponents do
  @moduledoc """
  Custom UI components for Farmec application.

  These components are converted from React and provide
  reusable UI elements throughout the application.
  """
  use Phoenix.Component

  # Helper function for PDF download paths
  defp pdf_path(type, id), do: "/api/pdf/#{type}/#{id}"

  @doc """
  Renders social media links.

  ## Examples

      <.social_links
        facebook="https://facebook.com/farmec"
        twitter="https://twitter.com/farmec"
      />
  """
  attr :facebook, :string, default: nil
  attr :twitter, :string, default: nil
  attr :instagram, :string, default: nil
  attr :linkedin, :string, default: nil
  attr :website, :string, default: nil
  attr :youtube, :string, default: nil
  attr :class, :string, default: ""

  def social_links(assigns) do
    ~H"""
    <div class={"social-links #{@class}"}>
      <%= if @facebook do %>
        <a href={@facebook} target="_blank" rel="noopener noreferrer" class="facebook-button">
          <i class="fa-brands fa-facebook"></i>
        </a>
      <% end %>

      <%= if @twitter do %>
        <a href={@twitter} target="_blank" rel="noopener noreferrer" class="twitter-button">
          <i class="fa-brands fa-twitter"></i>
        </a>
      <% end %>

      <%= if @instagram do %>
        <a href={@instagram} target="_blank" rel="noopener noreferrer" class="instagram-button">
          <i class="fa-brands fa-instagram"></i>
        </a>
      <% end %>

      <%= if @linkedin do %>
        <a href={@linkedin} target="_blank" rel="noopener noreferrer" class="linkedin-button">
          <i class="fa-brands fa-linkedin"></i>
        </a>
      <% end %>

      <%= if @website do %>
        <a href={@website} target="_blank" rel="noopener noreferrer" class="website-button">
          <i class="fa-solid fa-globe"></i>
        </a>
      <% end %>

      <%= if @youtube do %>
        <a href={@youtube} target="_blank" rel="noopener noreferrer" class="youtube-button">
          <i class="fa-brands fa-youtube"></i>
        </a>
      <% end %>
    </div>
    """
  end

  @doc """
  Renders a scroll-to-top button with smooth scrolling behavior.

  ## Examples

      <.to_top_button />
  """
  attr :class, :string, default: ""

  def to_top_button(assigns) do
    ~H"""
    <button
      id="toTopButton"
      aria-label="scroll-to-top-button"
      class={"to-top-button #{@class}"}
      phx-hook="ScrollToTop"
    >
      <i class="fa-solid fa-arrow-up"></i>
    </button>
    """
  end

  @doc """
  Renders a timeline card with title, date, and body content.

  ## Examples

      <.timeline_card
        timeline={%{id: "1", title: "Founded", date: "1999", body: "Company started"}}
        is_admin={true}
      />
  """
  attr :timeline, :map, required: true
  attr :is_admin, :boolean, default: false
  attr :class, :string, default: ""
  slot :inner_block

  def timeline_card(assigns) do
    ~H"""
    <div class={"timeline-card #{@class}"}>
      <h1 class="main-heading">{@timeline.title}</h1>
      <h2 class="timeline-date">
        <i class="fa-solid fa-clock"></i>
        {@timeline.date}
      </h2>
      <p class="paragraph">{@timeline.body}</p>

      <%= if @is_admin && @timeline.id do %>
        <div class="options-btn">
          {render_slot(@inner_block)}
        </div>
      <% end %>
    </div>
    """
  end

  @doc """
  Renders a delete button (requires LiveView context for functionality).

  ## Examples

      <.delete_button
        id="123"
        resource_key="timelines"
        phx-click="delete_resource"
      />
  """
  attr :id, :string, required: true
  attr :resource_key, :string, required: true
  attr :navigate_back, :boolean, default: false
  attr :rest, :global, include: ~w(phx-click phx-target)

  def delete_button(assigns) do
    ~H"""
    <button
      class="btn-form"
      phx-click="delete_resource"
      phx-value-id={@id}
      phx-value-resource={@resource_key}
      phx-value-navigate-back={@navigate_back}
      data-confirm="Are you sure you want to delete this item?"
      {@rest}
    >
      <i class="fa-solid fa-trash"></i>
    </button>
    """
  end

  @doc """
  Renders a download PDF button (requires server-side PDF generation).

  ## Examples

      <.download_pdf_button
        type="warranty"
        id="123"
      />
  """
  attr :type, :string, required: true, values: ~w(warranty registration)
  attr :id, :string, required: true
  attr :class, :string, default: "btn"

  def download_pdf_button(assigns) do
    ~H"""
    <a
      href={pdf_path(@type, @id)}
      download
      class={@class}
    >
      Download Form <i class="fa-solid fa-download"></i>
    </a>
    """
  end

  @doc """
  Renders a Google Map with a marker.

  Requires Google Maps API key to be loaded in the layout.
  Uses the GoogleMap Phoenix LiveView hook for initialization.

  ## Examples

      <.google_map
        lat={53.49200990196934}
        lng={-6.5423895598058435}
      />
  """
  attr :lat, :float, default: 53.49200990196934
  attr :lng, :float, default: -6.5423895598058435
  attr :width, :string, default: "600px"
  attr :height, :string, default: "600px"
  attr :class, :string, default: ""

  def google_map(assigns) do
    ~H"""
    <div class={"map-container #{@class}"}>
      <div
        id="google-map"
        phx-hook="GoogleMap"
        data-lat={@lat}
        data-lng={@lng}
        style={"width: #{@width}; height: #{@height};"}
      >
      </div>
    </div>
    """
  end

  @doc """
  Renders an image carousel with navigation controls.

  ## Examples

      <.carousel images={["img1.jpg", "img2.jpg"]} />
  """
  attr :images, :list, required: true
  attr :class, :string, default: ""

  def carousel(assigns) do
    ~H"""
    <section id="Hero" class={@class}>
      <div class="hero-container">
        <div class="slideshow" id="carousel" phx-hook="Carousel">
          <%= for {image, index} <- Enum.with_index(@images) do %>
            <img
              src={image}
              alt="Slide"
              class={"slides #{if index == 0, do: "fade-in", else: "fade-out"}"}
              data-index={index}
              onerror="this.src='/images/default.jpg'"
            />
          <% end %>
          <button class="prev-button" aria-label="previous slide">
            <i class="fa-solid fa-chevron-left"></i>
          </button>
          <button class="next-button" aria-label="next slide">
            <i class="fa-solid fa-chevron-right"></i>
          </button>
        </div>
      </div>
      <div class="typewriter">
        <h1>Importers & Distributors of Quality Agricultural Machinery</h1>
        <a href="#Info" class="btn">
          Find Out More: <i class="fa-solid fa-chevron-circle-down"></i>
        </a>
      </div>
    </section>
    """
  end

  @doc """
  Renders the displays section with stats and specials.

  ## Examples

      <.displays_section stats={stats} specials={specials} />
  """
  attr :stats, :list, required: true
  attr :specials, :list, required: true
  attr :class, :string, default: ""

  def displays_section(assigns) do
    ~H"""
    <section id="Info" class={@class}>
      <div class="info-section">
        <h1 class="section-heading">Farmec At A Glance:</h1>
        <p class="sub-heading">
          This is a Quick Look at what Separates us from our Competitors
        </p>
        <div class="stats">
          <%= for item <- @stats do %>
            <a href={item.link}>
              <ul class="stat-list">
                <li class="stat-list-item"><%= item.title %></li>
                <li class="stat-list-item"><i class={item.icon}></i></li>
                <li class="stat-list-item"><%= item.description %></li>
              </ul>
            </a>
          <% end %>
        </div>
      </div>
      <div class="info-section">
        <h1 class="section-heading">What Can We Offer:</h1>
        <p class="sub-heading">
          Farmec is committed to its customers and guarantees the following
        </p>
        <div class="specials">
          <%= for item <- @specials do %>
            <a href={item.link}>
              <ul class="special-list">
                <li class="special-list-item"><%= item.title %></li>
                <li class="special-list-item"><i class={item.icon}></i></li>
                <li class="special-list-item"><%= item.description %></li>
              </ul>
            </a>
          <% end %>
        </div>
      </div>
    </section>
    """
  end

  @doc """
  Renders the contact section with form, map, and business info.

  ## Examples

      <.contact_section />
  """
  attr :class, :string, default: ""
  slot :inner_block

  def contact_section(assigns) do
    ~H"""
    <section id="contact" class={@class}>
      <h1 class="section-heading">Contact Us:</h1>
      <div class="contact-section">
        <%= render_slot(@inner_block) %>

        <.google_map />

        <div class="info-section">
          <h1 class="sub-heading">Business Information:</h1>
          <div class="info">
            <div class="info-item">
              Opening Hours:
              <br />
              <span class="info-item-text">Monday - Friday: 9am - 5:30pm</span>
            </div>
            <div class="info-item">
              Telephone:
              <br />
              <span class="info-item-text">
                <a href="tel:018259289">01 825 9289</a>
              </span>
            </div>
            <div class="info-item">
              International:
              <br />
              <span class="info-item-text">
                <a href="tel:+35318259289">+353 1 825 9289</a>
              </span>
            </div>
            <div class="info-item">
              Email:
              <br />
              <span class="info-item-text">Info@farmec.ie</span>
            </div>
            <div class="info-item">
              Address:
              <br />
              <span class="info-item-text">Clonross, Drumree, Co. Meath, A85PK30</span>
            </div>
            <div class="info-item">
              <div class="social-links">
                <a
                  class="socials"
                  href="https://www.facebook.com/FarmecIreland/"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="Visit our Facebook page"
                >
                  <i class="fa-brands fa-facebook"></i>
                </a>
                <a
                  class="socials"
                  href="https://twitter.com/farmec1?lang=en"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="Visit our Twitter page"
                >
                  <i class="fa-brands fa-twitter"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    """
  end
end

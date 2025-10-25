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
end

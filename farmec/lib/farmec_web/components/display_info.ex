defmodule FarmecWeb.DisplayInfo do
  @moduledoc """
  Static display information and configuration data for the Farmec website.
  Converted from React displaysInfo.tsx
  """

  @doc """
  Returns the list of special features/items for the home page.
  """
  def specials_items do
    [
      %{
        title: "Quality Stock",
        description:
          "Farmec is committed to the importation and distribution of only quality brands of unique farm machinery. We guarantee that all our suppliers are committed to providing farmers with durable and superior stock",
        icon: "fa-solid fa-tractor",
        link: "/suppliers"
      },
      %{
        title: "Assembly",
        description:
          "Farmec have a team of qualified and experienced staff that ensure abundant care is taken during the assembly process; we make sure that a quality supply chain is maintained from manufacturer to customer",
        icon: "fa-solid fa-toolbox",
        link: "/suppliers"
      },
      %{
        title: "Spare Parts",
        description:
          "Farmec offers a diverse and complete range of spare parts for all its machinery. Quality stock control and industry expertise ensures parts finds their way to you efficiently",
        icon: "fa-solid fa-gears",
        link: "/spareparts"
      },
      %{
        title: "Customer Service",
        description:
          "Farmec is a family run company, we make sure we extend the ethos of a small community to our customers. We build established relationships with our dealers that provide them and the farmers with extensive guidance",
        icon: "fa-solid fa-user-plus",
        link: "/contact"
      }
    ]
  end

  @doc """
  Returns the list of statistics items for the home page.
  """
  def stats_items do
    [
      %{
        title: "Large Network",
        description: "50+ Dealers Nationwide",
        icon: "fa-solid fa-users",
        link: "/suppliers"
      },
      %{
        title: "Experience",
        description: "25+ Years in Business",
        icon: "fa-solid fa-business-time",
        link: "/about"
      },
      %{
        title: "Diverse Range",
        description: "10+ Quality Suppliers",
        icon: "fa-solid fa-handshake",
        link: "/suppliers"
      },
      %{
        title: "Commitment",
        description: "Warranty Guarantee",
        icon: "fa-solid fa-wrench",
        link: "/spareparts"
      }
    ]
  end

  @doc """
  Helper component to render a special/stat item card.

  ## Examples

      <%= for item <- FarmecWeb.DisplayInfo.specials_items() do %>
        <.info_card item={item} />
      <% end %>
  """
  use Phoenix.Component

  attr :item, :map, required: true
  attr :class, :string, default: ""

  def info_card(assigns) do
    ~H"""
    <div class={"info-card #{@class}"}>
      <i class={@item.icon}></i>
      <h3><%= @item.title %></h3>
      <p><%= @item.description %></p>
      <a href={@item.link} class="info-link">Learn More</a>
    </div>
    """
  end
end

defmodule FarmecWeb.PageHTML do
  @moduledoc """
  This module contains pages rendered by PageController.

  See the `page_html` directory for all templates available.
  """
  use FarmecWeb, :html

  # Import our custom components
  import FarmecWeb.CustomComponents

  embed_templates "page_html/*"
end

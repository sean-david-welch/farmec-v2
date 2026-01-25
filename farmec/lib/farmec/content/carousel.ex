defmodule Farmec.Content.Carousel do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Carousel" do
    field :name, :string
    field :image, :string
    field :created, :string
  end

  @doc false
  def changeset(carousel, attrs) do
    carousel
    |> cast(attrs, [:name, :image, :created])
    |> validate_required([:name])
    |> validate_length(:name, min: 2, max: 255)
  end
end

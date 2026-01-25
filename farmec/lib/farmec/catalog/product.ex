defmodule Farmec.Catalog.Product do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Product" do
    field :machine_id, :string
    field :name, :string
    field :product_image, :string
    field :description, :string
    field :product_link, :string

    belongs_to :machine, Farmec.Catalog.Machine, define_field: false, foreign_key: :machine_id, type: :string
  end

  @doc false
  def changeset(product, attrs) do
    product
    |> cast(attrs, [:machine_id, :name, :product_image, :description, :product_link])
    |> validate_required([:machine_id, :name])
    |> validate_length(:name, min: 2, max: 255)
    |> foreign_key_constraint(:machine_id)
  end
end

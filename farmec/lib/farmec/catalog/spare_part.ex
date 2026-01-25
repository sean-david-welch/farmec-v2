defmodule Farmec.Catalog.SparePart do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "SpareParts" do
    field :supplier_id, :string
    field :name, :string
    field :parts_image, :string
    field :spare_parts_link, :string

    belongs_to :supplier, Farmec.Catalog.Supplier, define_field: false, foreign_key: :supplier_id, type: :string
  end

  @doc false
  def changeset(spare_part, attrs) do
    spare_part
    |> cast(attrs, [:supplier_id, :name, :parts_image, :spare_parts_link])
    |> validate_required([:supplier_id, :name])
    |> validate_length(:name, min: 2, max: 255)
    |> foreign_key_constraint(:supplier_id)
  end
end

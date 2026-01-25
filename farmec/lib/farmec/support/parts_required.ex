defmodule Farmec.Support.PartsRequired do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "PartsRequired" do
    field :warranty_id, :string
    field :part_number, :string
    field :quantity_needed, :string
    field :invoice_number, :string
    field :description, :string

    belongs_to :warranty_claim, Farmec.Support.WarrantyClaim, define_field: false, foreign_key: :warranty_id, type: :string
  end

  @doc false
  def changeset(parts_required, attrs) do
    parts_required
    |> cast(attrs, [:warranty_id, :part_number, :quantity_needed, :invoice_number, :description])
    |> validate_required([:warranty_id, :quantity_needed])
    |> foreign_key_constraint(:warranty_id)
  end
end

defmodule Farmec.Catalog.Machine do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Machine" do
    field :name, :string
    field :machine_image, :string
    field :description, :string
    field :machine_link, :string
    field :slug, :string
    field :created, :string

    belongs_to :supplier, Farmec.Catalog.Supplier, foreign_key: :supplier_id, type: :string
  end

  @doc false
  def changeset(machine, attrs) do
    machine
    |> cast(attrs, [:name, :machine_image, :description, :machine_link, :slug, :supplier_id])
    |> validate_required([:name, :supplier_id])
    |> validate_length(:name, min: 2, max: 255)
    |> foreign_key_constraint(:supplier_id)
  end
end

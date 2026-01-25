defmodule Farmec.Support.WarrantyClaim do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "WarrantyClaim" do
    field :dealer, :string
    field :dealer_contact, :string
    field :owner_name, :string
    field :machine_model, :string
    field :serial_number, :string
    field :install_date, :string
    field :failure_date, :string
    field :repair_date, :string
    field :failure_details, :string
    field :repair_details, :string
    field :labour_hours, :string
    field :completed_by, :string
    field :created, :string

    has_many :parts_required, Farmec.Support.PartsRequired, foreign_key: :warranty_id
  end

  @doc false
  def changeset(warranty_claim, attrs) do
    warranty_claim
    |> cast(attrs, [
      :dealer,
      :dealer_contact,
      :owner_name,
      :machine_model,
      :serial_number,
      :install_date,
      :failure_date,
      :repair_date,
      :failure_details,
      :repair_details,
      :labour_hours,
      :completed_by,
      :created
    ])
    |> validate_required([:dealer, :owner_name, :machine_model, :serial_number])
  end
end

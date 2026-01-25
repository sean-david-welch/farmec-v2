defmodule Farmec.Support.MachineRegistration do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "MachineRegistration" do
    field :dealer_name, :string
    field :dealer_address, :string
    field :owner_name, :string
    field :owner_address, :string
    field :machine_model, :string
    field :serial_number, :string
    field :install_date, :string
    field :invoice_number, :string
    field :complete_supply, :integer
    field :pdi_complete, :integer
    field :pto_correct, :integer
    field :machine_test_run, :integer
    field :safety_induction, :integer
    field :operator_handbook, :integer
    field :date, :string
    field :completed_by, :string
    field :created, :string
  end

  @doc false
  def changeset(machine_registration, attrs) do
    machine_registration
    |> cast(attrs, [
      :dealer_name,
      :dealer_address,
      :owner_name,
      :owner_address,
      :machine_model,
      :serial_number,
      :install_date,
      :invoice_number,
      :complete_supply,
      :pdi_complete,
      :pto_correct,
      :machine_test_run,
      :safety_induction,
      :operator_handbook,
      :date,
      :completed_by,
      :created
    ])
    |> validate_required([:dealer_name, :owner_name, :machine_model, :serial_number])
    |> validate_inclusion(:complete_supply, [0, 1])
    |> validate_inclusion(:pdi_complete, [0, 1])
    |> validate_inclusion(:pto_correct, [0, 1])
    |> validate_inclusion(:machine_test_run, [0, 1])
    |> validate_inclusion(:safety_induction, [0, 1])
    |> validate_inclusion(:operator_handbook, [0, 1])
  end
end

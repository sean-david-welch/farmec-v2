defmodule Farmec.Team.Employee do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Employee" do
    field :name, :string
    field :email, :string
    field :role, :string
    field :profile_image, :string
    field :created, :string
  end

  @doc false
  def changeset(employee, attrs) do
    employee
    |> cast(attrs, [:name, :email, :role, :profile_image])
    |> validate_required([:name, :email, :role])
    |> validate_format(:email, ~r/^[^\s]+@[^\s]+$/, message: "must be a valid email")
    |> validate_length(:name, min: 2, max: 255)
  end
end

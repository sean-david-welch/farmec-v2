defmodule Farmec.Content.Exhibition do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Exhibition" do
    field :title, :string
    field :date, :string
    field :location, :string
    field :info, :string
    field :created, :string
  end

  @doc false
  def changeset(exhibition, attrs) do
    exhibition
    |> cast(attrs, [:title, :date, :location, :info, :created])
    |> validate_required([:title])
    |> validate_length(:title, min: 2, max: 500)
  end
end

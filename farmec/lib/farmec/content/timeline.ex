defmodule Farmec.Content.Timeline do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Timeline" do
    field :title, :string
    field :date, :string
    field :body, :string
    field :created, :string
  end

  @doc false
  def changeset(timeline, attrs) do
    timeline
    |> cast(attrs, [:title, :date, :body, :created])
    |> validate_required([:title])
    |> validate_length(:title, min: 2, max: 500)
  end
end

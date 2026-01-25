defmodule Farmec.Content.Video do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Video" do
    field :supplier_id, :string
    field :web_url, :string
    field :title, :string
    field :description, :string
    field :video_id, :string
    field :thumbnail_url, :string
    field :created, :string

    belongs_to :supplier, Farmec.Catalog.Supplier, define_field: false, foreign_key: :supplier_id, type: :string
  end

  @doc false
  def changeset(video, attrs) do
    video
    |> cast(attrs, [:supplier_id, :web_url, :title, :description, :video_id, :thumbnail_url, :created])
    |> validate_required([:supplier_id])
    |> foreign_key_constraint(:supplier_id)
  end
end

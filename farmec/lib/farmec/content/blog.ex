defmodule Farmec.Content.Blog do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Blog" do
    field :title, :string
    field :date, :string
    field :main_image, :string
    field :subheading, :string
    field :body, :string
    field :created, :string
  end

  @doc false
  def changeset(blog, attrs) do
    blog
    |> cast(attrs, [:title, :date, :main_image, :subheading, :body, :created])
    |> validate_required([:title])
    |> validate_length(:title, min: 2, max: 500)
  end
end

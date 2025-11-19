defmodule Farmec.Catalog.Supplier do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :string, autogenerate: false}
  @foreign_key_type :string

  schema "Supplier" do
    field :name, :string
    field :logo_image, :string
    field :marketing_image, :string
    field :description, :string
    field :social_facebook, :string
    field :social_twitter, :string
    field :social_instagram, :string
    field :social_youtube, :string
    field :social_linkedin, :string
    field :social_website, :string
    field :slug, :string
    field :created, :string

    has_many :machines, Farmec.Catalog.Machine, foreign_key: :supplier_id
  end

  @doc false
  def changeset(supplier, attrs) do
    supplier
    |> cast(attrs, [
      :name,
      :logo_image,
      :marketing_image,
      :description,
      :social_facebook,
      :social_twitter,
      :social_instagram,
      :social_youtube,
      :social_linkedin,
      :social_website,
      :slug
    ])
    |> validate_required([:name])
    |> validate_length(:name, min: 2, max: 255)
  end
end

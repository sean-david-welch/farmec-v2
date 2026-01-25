defmodule Farmec.Content do
  @moduledoc """
  The Content context for managing blogs, carousels, exhibitions, timelines, and videos.
  """

  import Ecto.Query, warn: false
  alias Farmec.Repo
  alias Farmec.Content.{Blog, Carousel, Exhibition, Timeline, Video}

  ## Blogs

  @doc """
  Returns the list of blogs.
  """
  def list_blogs do
    Repo.all(Blog)
  end

  @doc """
  Gets a single blog.
  Raises `Ecto.NoResultsError` if the Blog does not exist.
  """
  def get_blog!(id), do: Repo.get!(Blog, id)

  @doc """
  Creates a blog.
  """
  def create_blog(attrs \\ %{}) do
    %Blog{}
    |> Blog.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a blog.
  """
  def update_blog(%Blog{} = blog, attrs) do
    blog
    |> Blog.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a blog.
  """
  def delete_blog(%Blog{} = blog) do
    Repo.delete(blog)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking blog changes.
  """
  def change_blog(%Blog{} = blog, attrs \\ %{}) do
    Blog.changeset(blog, attrs)
  end

  ## Carousels

  @doc """
  Returns the list of carousels.
  """
  def list_carousels do
    Repo.all(Carousel)
  end

  @doc """
  Gets a single carousel.
  Raises `Ecto.NoResultsError` if the Carousel does not exist.
  """
  def get_carousel!(id), do: Repo.get!(Carousel, id)

  @doc """
  Creates a carousel.
  """
  def create_carousel(attrs \\ %{}) do
    %Carousel{}
    |> Carousel.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a carousel.
  """
  def update_carousel(%Carousel{} = carousel, attrs) do
    carousel
    |> Carousel.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a carousel.
  """
  def delete_carousel(%Carousel{} = carousel) do
    Repo.delete(carousel)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking carousel changes.
  """
  def change_carousel(%Carousel{} = carousel, attrs \\ %{}) do
    Carousel.changeset(carousel, attrs)
  end

  ## Exhibitions

  @doc """
  Returns the list of exhibitions.
  """
  def list_exhibitions do
    Repo.all(Exhibition)
  end

  @doc """
  Gets a single exhibition.
  Raises `Ecto.NoResultsError` if the Exhibition does not exist.
  """
  def get_exhibition!(id), do: Repo.get!(Exhibition, id)

  @doc """
  Creates an exhibition.
  """
  def create_exhibition(attrs \\ %{}) do
    %Exhibition{}
    |> Exhibition.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates an exhibition.
  """
  def update_exhibition(%Exhibition{} = exhibition, attrs) do
    exhibition
    |> Exhibition.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes an exhibition.
  """
  def delete_exhibition(%Exhibition{} = exhibition) do
    Repo.delete(exhibition)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking exhibition changes.
  """
  def change_exhibition(%Exhibition{} = exhibition, attrs \\ %{}) do
    Exhibition.changeset(exhibition, attrs)
  end

  ## Timelines

  @doc """
  Returns the list of timelines.
  """
  def list_timelines do
    Repo.all(Timeline)
  end

  @doc """
  Gets a single timeline.
  Raises `Ecto.NoResultsError` if the Timeline does not exist.
  """
  def get_timeline!(id), do: Repo.get!(Timeline, id)

  @doc """
  Creates a timeline.
  """
  def create_timeline(attrs \\ %{}) do
    %Timeline{}
    |> Timeline.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a timeline.
  """
  def update_timeline(%Timeline{} = timeline, attrs) do
    timeline
    |> Timeline.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a timeline.
  """
  def delete_timeline(%Timeline{} = timeline) do
    Repo.delete(timeline)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking timeline changes.
  """
  def change_timeline(%Timeline{} = timeline, attrs \\ %{}) do
    Timeline.changeset(timeline, attrs)
  end

  ## Videos

  @doc """
  Returns the list of videos.
  """
  def list_videos do
    Repo.all(Video)
  end

  @doc """
  Returns videos for a specific supplier.
  """
  def list_videos_by_supplier(supplier_id) do
    Video
    |> where([v], v.supplier_id == ^supplier_id)
    |> Repo.all()
  end

  @doc """
  Gets a single video.
  Raises `Ecto.NoResultsError` if the Video does not exist.
  """
  def get_video!(id), do: Repo.get!(Video, id)

  @doc """
  Gets a video with preloaded supplier.
  """
  def get_video_with_supplier(id) do
    Video
    |> Repo.get!(id)
    |> Repo.preload(:supplier)
  end

  @doc """
  Creates a video.
  """
  def create_video(attrs \\ %{}) do
    %Video{}
    |> Video.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a video.
  """
  def update_video(%Video{} = video, attrs) do
    video
    |> Video.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a video.
  """
  def delete_video(%Video{} = video) do
    Repo.delete(video)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking video changes.
  """
  def change_video(%Video{} = video, attrs \\ %{}) do
    Video.changeset(video, attrs)
  end
end

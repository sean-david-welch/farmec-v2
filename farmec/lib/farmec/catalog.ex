defmodule Farmec.Catalog do
  @moduledoc """
  The Catalog context for managing suppliers, machines, products, and spare parts.
  """

  import Ecto.Query, warn: false
  alias Farmec.Repo
  alias Farmec.Catalog.{Supplier, Machine, Product, SparePart}

  ## Suppliers

  @doc """
  Returns the list of suppliers.

  ## Examples

      iex> list_suppliers()
      [%Supplier{}, ...]

  """
  def list_suppliers do
    Repo.all(Supplier)
  end

  @doc """
  Gets a single supplier.

  Raises `Ecto.NoResultsError` if the Supplier does not exist.

  ## Examples

      iex> get_supplier!(123)
      %Supplier{}

      iex> get_supplier!(456)
      ** (Ecto.NoResultsError)

  """
  def get_supplier!(id), do: Repo.get!(Supplier, id)

  @doc """
  Gets a supplier with preloaded machines.
  """
  def get_supplier_with_machines(id) do
    Supplier
    |> Repo.get!(id)
    |> Repo.preload(:machines)
  end

  @doc """
  Creates a supplier.

  ## Examples

      iex> create_supplier(%{field: value})
      {:ok, %Supplier{}}

      iex> create_supplier(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_supplier(attrs \\ %{}) do
    %Supplier{}
    |> Supplier.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a supplier.

  ## Examples

      iex> update_supplier(supplier, %{field: new_value})
      {:ok, %Supplier{}}

      iex> update_supplier(supplier, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_supplier(%Supplier{} = supplier, attrs) do
    supplier
    |> Supplier.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a supplier.

  ## Examples

      iex> delete_supplier(supplier)
      {:ok, %Supplier{}}

      iex> delete_supplier(supplier)
      {:error, %Ecto.Changeset{}}

  """
  def delete_supplier(%Supplier{} = supplier) do
    Repo.delete(supplier)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking supplier changes.

  ## Examples

      iex> change_supplier(supplier)
      %Ecto.Changeset{data: %Supplier{}}

  """
  def change_supplier(%Supplier{} = supplier, attrs \\ %{}) do
    Supplier.changeset(supplier, attrs)
  end

  ## Machines

  @doc """
  Returns the list of machines.

  ## Examples

      iex> list_machines()
      [%Machine{}, ...]

  """
  def list_machines do
    Repo.all(Machine)
  end

  @doc """
  Returns machines for a specific supplier.
  """
  def list_machines_by_supplier(supplier_id) do
    Machine
    |> where([m], m.supplier_id == ^supplier_id)
    |> Repo.all()
  end

  @doc """
  Gets a single machine.

  Raises `Ecto.NoResultsError` if the Machine does not exist.

  ## Examples

      iex> get_machine!(123)
      %Machine{}

      iex> get_machine!(456)
      ** (Ecto.NoResultsError)

  """
  def get_machine!(id), do: Repo.get!(Machine, id)

  @doc """
  Gets a machine with preloaded supplier.
  """
  def get_machine_with_supplier(id) do
    Machine
    |> Repo.get!(id)
    |> Repo.preload(:supplier)
  end

  @doc """
  Creates a machine.

  ## Examples

      iex> create_machine(%{field: value})
      {:ok, %Machine{}}

      iex> create_machine(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_machine(attrs \\ %{}) do
    %Machine{}
    |> Machine.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a machine.

  ## Examples

      iex> update_machine(machine, %{field: new_value})
      {:ok, %Machine{}}

      iex> update_machine(machine, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_machine(%Machine{} = machine, attrs) do
    machine
    |> Machine.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a machine.

  ## Examples

      iex> delete_machine(machine)
      {:ok, %Machine{}}

      iex> delete_machine(machine)
      {:error, %Ecto.Changeset{}}

  """
  def delete_machine(%Machine{} = machine) do
    Repo.delete(machine)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking machine changes.

  ## Examples

      iex> change_machine(machine)
      %Ecto.Changeset{data: %Machine{}}

  """
  def change_machine(%Machine{} = machine, attrs \\ %{}) do
    Machine.changeset(machine, attrs)
  end

  ## Products

  @doc """
  Returns the list of products.
  """
  def list_products do
    Repo.all(Product)
  end

  @doc """
  Returns products for a specific machine.
  """
  def list_products_by_machine(machine_id) do
    Product
    |> where([p], p.machine_id == ^machine_id)
    |> Repo.all()
  end

  @doc """
  Gets a single product.
  Raises `Ecto.NoResultsError` if the Product does not exist.
  """
  def get_product!(id), do: Repo.get!(Product, id)

  @doc """
  Gets a product with preloaded machine.
  """
  def get_product_with_machine(id) do
    Product
    |> Repo.get!(id)
    |> Repo.preload(:machine)
  end

  @doc """
  Creates a product.
  """
  def create_product(attrs \\ %{}) do
    %Product{}
    |> Product.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a product.
  """
  def update_product(%Product{} = product, attrs) do
    product
    |> Product.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a product.
  """
  def delete_product(%Product{} = product) do
    Repo.delete(product)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking product changes.
  """
  def change_product(%Product{} = product, attrs \\ %{}) do
    Product.changeset(product, attrs)
  end

  ## Spare Parts

  @doc """
  Returns the list of spare parts.
  """
  def list_spare_parts do
    Repo.all(SparePart)
  end

  @doc """
  Returns spare parts for a specific supplier.
  """
  def list_spare_parts_by_supplier(supplier_id) do
    SparePart
    |> where([sp], sp.supplier_id == ^supplier_id)
    |> Repo.all()
  end

  @doc """
  Gets a single spare part.
  Raises `Ecto.NoResultsError` if the SparePart does not exist.
  """
  def get_spare_part!(id), do: Repo.get!(SparePart, id)

  @doc """
  Gets a spare part with preloaded supplier.
  """
  def get_spare_part_with_supplier(id) do
    SparePart
    |> Repo.get!(id)
    |> Repo.preload(:supplier)
  end

  @doc """
  Creates a spare part.
  """
  def create_spare_part(attrs \\ %{}) do
    %SparePart{}
    |> SparePart.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a spare part.
  """
  def update_spare_part(%SparePart{} = spare_part, attrs) do
    spare_part
    |> SparePart.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a spare part.
  """
  def delete_spare_part(%SparePart{} = spare_part) do
    Repo.delete(spare_part)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking spare part changes.
  """
  def change_spare_part(%SparePart{} = spare_part, attrs \\ %{}) do
    SparePart.changeset(spare_part, attrs)
  end
end

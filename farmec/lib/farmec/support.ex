defmodule Farmec.Support do
  @moduledoc """
  The Support context for managing warranty claims, machine registrations, and parts required.
  """

  import Ecto.Query, warn: false
  alias Farmec.Repo
  alias Farmec.Support.{WarrantyClaim, MachineRegistration, PartsRequired}

  ## Warranty Claims

  @doc """
  Returns the list of warranty claims.
  """
  def list_warranty_claims do
    Repo.all(WarrantyClaim)
  end

  @doc """
  Gets a single warranty claim.
  Raises `Ecto.NoResultsError` if the WarrantyClaim does not exist.
  """
  def get_warranty_claim!(id), do: Repo.get!(WarrantyClaim, id)

  @doc """
  Gets a warranty claim with preloaded parts required.
  """
  def get_warranty_claim_with_parts(id) do
    WarrantyClaim
    |> Repo.get!(id)
    |> Repo.preload(:parts_required)
  end

  @doc """
  Creates a warranty claim.
  """
  def create_warranty_claim(attrs \\ %{}) do
    %WarrantyClaim{}
    |> WarrantyClaim.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a warranty claim.
  """
  def update_warranty_claim(%WarrantyClaim{} = warranty_claim, attrs) do
    warranty_claim
    |> WarrantyClaim.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a warranty claim.
  """
  def delete_warranty_claim(%WarrantyClaim{} = warranty_claim) do
    Repo.delete(warranty_claim)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking warranty claim changes.
  """
  def change_warranty_claim(%WarrantyClaim{} = warranty_claim, attrs \\ %{}) do
    WarrantyClaim.changeset(warranty_claim, attrs)
  end

  ## Machine Registrations

  @doc """
  Returns the list of machine registrations.
  """
  def list_machine_registrations do
    Repo.all(MachineRegistration)
  end

  @doc """
  Gets a single machine registration.
  Raises `Ecto.NoResultsError` if the MachineRegistration does not exist.
  """
  def get_machine_registration!(id), do: Repo.get!(MachineRegistration, id)

  @doc """
  Creates a machine registration.
  """
  def create_machine_registration(attrs \\ %{}) do
    %MachineRegistration{}
    |> MachineRegistration.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a machine registration.
  """
  def update_machine_registration(%MachineRegistration{} = machine_registration, attrs) do
    machine_registration
    |> MachineRegistration.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a machine registration.
  """
  def delete_machine_registration(%MachineRegistration{} = machine_registration) do
    Repo.delete(machine_registration)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking machine registration changes.
  """
  def change_machine_registration(%MachineRegistration{} = machine_registration, attrs \\ %{}) do
    MachineRegistration.changeset(machine_registration, attrs)
  end

  ## Parts Required

  @doc """
  Returns the list of parts required.
  """
  def list_parts_required do
    Repo.all(PartsRequired)
  end

  @doc """
  Returns parts required for a specific warranty claim.
  """
  def list_parts_required_by_warranty(warranty_id) do
    PartsRequired
    |> where([pr], pr.warranty_id == ^warranty_id)
    |> Repo.all()
  end

  @doc """
  Gets a single parts required.
  Raises `Ecto.NoResultsError` if the PartsRequired does not exist.
  """
  def get_parts_required!(id), do: Repo.get!(PartsRequired, id)

  @doc """
  Creates a parts required.
  """
  def create_parts_required(attrs \\ %{}) do
    %PartsRequired{}
    |> PartsRequired.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a parts required.
  """
  def update_parts_required(%PartsRequired{} = parts_required, attrs) do
    parts_required
    |> PartsRequired.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a parts required.
  """
  def delete_parts_required(%PartsRequired{} = parts_required) do
    Repo.delete(parts_required)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking parts required changes.
  """
  def change_parts_required(%PartsRequired{} = parts_required, attrs \\ %{}) do
    PartsRequired.changeset(parts_required, attrs)
  end
end

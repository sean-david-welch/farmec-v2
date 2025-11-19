defmodule Farmec.Repo.Migrations.BaselineSchema do
  use Ecto.Migration

  def change do
    # This is a baseline migration for an existing SQLite database.
    # The tables already exist in the database copied from production.
    # This migration just marks the schema as being at this state.
    # Future migrations will build on top of this baseline.

    # No operations needed - tables already exist
    :ok
  end
end

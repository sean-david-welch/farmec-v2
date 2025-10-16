defmodule Farmec.Repo do
  use Ecto.Repo,
    otp_app: :farmec,
    adapter: Ecto.Adapters.SQLite3
end

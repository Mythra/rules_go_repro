load("@bazel_skylib//rules:common_settings.bzl", "string_flag")

string_flag(
  name = 'sql_driver',
  values = ['sqlite', 'postgres'],
  build_setting_default = 'sqlite',
)
config_setting(
  name = 'sql_driver_sqlite',
  flag_values = {
    ':sql_driver': 'sqlite',
  }
)
config_setting(
  name = 'sql_driver_postgres',
  flag_values = {
    ':sql_driver': 'postgres',
  }
)
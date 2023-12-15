from generated.builders.dashboard import (
    Dashboard,
    TimePicker,
    Row,
    DatasourceVariable,
    QueryVariable,
)
from generated.models.dashboard import (
    DashboardCursorSync,
    VariableHide,
    VariableOption,
    VariableRefresh,
    DataSourceRef,
    VariableSort,
)
from generated.cog.encoder import JSONEncoder
from examples.python.raspberry.cpu import cpu_usage_timeseries, cpu_load_average_timeseries, cpu_temperature_gauge
from examples.python.raspberry.memory import memory_usage_timeseries, memory_usage_gauge


def build_dashboard() -> Dashboard:
    builder = (
        Dashboard("[TEST] Node Exporter / Raspberry")
        .uid("test-dashboard-raspberry")
        .tags(["generated", "raspberrypi-node-integration"])
        .refresh("30s")
        .time("now-30m", "now")
        #.timezone(TimeZoneBrowser)
        .timezone("browser")
        .timepicker(
            TimePicker()
            .refresh_intervals(["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"])
            .time_options(["5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"])
        )
        .tooltip(DashboardCursorSync.CROSSHAIR)
        # "Data Source" variable
        .with_variable(
            DatasourceVariable("datasource")
            .label("Data Source")
            .hide(VariableHide.DONT_HIDE)
            .type_val("prometheus")
            .current(VariableOption(selected=True, text="grafanacloud-potatopi-prom", value="grafanacloud-prom"))
        )
        # "Instance" variable
        .with_variable(
            QueryVariable("instance")
            .label("Instance")
            .hide(VariableHide.DONT_HIDE)
            .refresh(VariableRefresh.ON_TIME_RANGE_CHANGED)
            .query('label_values(node_uname_info{job="integrations/raspberrypi-node", sysname!="Darwin"}, instance)')
            .datasource(DataSourceRef(type_val="prometheus", uid="$datasource"))
            .current(VariableOption(selected=False, text="potato", value="potato"))
            .sort(VariableSort.DISABLED)
        )
        # CPU
        .with_row(Row("CPU"))
        .with_panel(cpu_usage_timeseries())
        .with_panel(cpu_temperature_gauge())
        .with_panel(cpu_load_average_timeseries())
        # Memory
        .with_row(Row("Memory"))
        .with_panel(memory_usage_timeseries())
        .with_panel(memory_usage_gauge())
        # Disk
        .with_row(Row("Disk"))
        # Network
        .with_row(Row("Network"))
        # Logs
        .with_row(Row("Logs"))
    )

    return builder


if __name__ == '__main__':
    dashboard = build_dashboard().build()
    encoder = JSONEncoder(sort_keys=True, indent=2)

    print(
        encoder.encode(dashboard)
    )
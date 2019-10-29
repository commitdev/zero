{
    "agent": {
        "region": "${region}"
    },
    "logs": {
        "metrics_collected": {
            "kubernetes": {
                "cluster_name": "${cluster_name}",
                "metrics_collection_interval": 60
            }
        },
        "force_flush_interval": 5
    },
    "metrics": {
        "metrics_collected": {
            "statsd": {
                "service_address": ":8125"
            }
        }
    }
}
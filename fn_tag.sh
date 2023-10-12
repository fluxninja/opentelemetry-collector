#!/usr/bin/env bash

set -uo pipefail

tag_components=(
    "component/"
    "confmap/"
    "consumer/"
    "exporter/loggingexporter/"
    "exporter/otlpexporter/"
    "exporter/otlphttpexporter/"
    "extension/ballastextension/"
    "extension/zpagesextension/"
    "pdata/"
    "processor/batchprocessor/"
    "processor/memorylimiterprocessor/"
    "receiver/otlpreceiver/"
    "connector/"
    "exporter/"
    "extension/"
    "otelcol/"
    "processor/"
    "receiver/"
    ""
)
upstream_version=$(curl -sL https://api.github.com/repos/open-telemetry/opentelemetry-collector/releases/latest | jq -r ".tag_name" -e | sed 's/cmd\/builder\///g')
otel_version="${2:-${upstream_version}}"
patch_number="${3:-1}"
tag_version="${otel_version}-fn.patch.${patch_number}"

tag_names=()
for component in "${tag_components[@]}"; do
    tag_name="${component}${tag_version}"
    tag_names+=($tag_name)
done

make_tags() {
    for tag_name in "${tag_names[@]}"; do
        echo "Tagging as ${tag_name}"
        git tag ${tag_name}
    done
    return
}

push_tags() {
    echo "Pushing tags ${tag_names[@]}"
    git push origin ${tag_names[@]}
    return
}

case "${1}" in
    make) make_tags;;
    push) push_tags;;
    *) printf 'Unknown command - "%s"\n' "${1}"; exit 1;;
esac

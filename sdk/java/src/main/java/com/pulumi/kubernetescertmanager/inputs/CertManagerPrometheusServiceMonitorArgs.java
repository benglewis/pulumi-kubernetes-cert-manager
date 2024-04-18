// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescertmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertManagerPrometheusServiceMonitorArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertManagerPrometheusServiceMonitorArgs Empty = new CertManagerPrometheusServiceMonitorArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="interval")
    private @Nullable Output<String> interval;

    public Optional<Output<String>> interval() {
        return Optional.ofNullable(this.interval);
    }

    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    @Import(name="path")
    private @Nullable Output<String> path;

    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    @Import(name="prometheusInstance")
    private @Nullable Output<String> prometheusInstance;

    public Optional<Output<String>> prometheusInstance() {
        return Optional.ofNullable(this.prometheusInstance);
    }

    @Import(name="string")
    private @Nullable Output<String> string;

    public Optional<Output<String>> string() {
        return Optional.ofNullable(this.string);
    }

    @Import(name="targetPort")
    private @Nullable Output<Integer> targetPort;

    public Optional<Output<Integer>> targetPort() {
        return Optional.ofNullable(this.targetPort);
    }

    private CertManagerPrometheusServiceMonitorArgs() {}

    private CertManagerPrometheusServiceMonitorArgs(CertManagerPrometheusServiceMonitorArgs $) {
        this.enabled = $.enabled;
        this.interval = $.interval;
        this.labels = $.labels;
        this.path = $.path;
        this.prometheusInstance = $.prometheusInstance;
        this.string = $.string;
        this.targetPort = $.targetPort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertManagerPrometheusServiceMonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertManagerPrometheusServiceMonitorArgs $;

        public Builder() {
            $ = new CertManagerPrometheusServiceMonitorArgs();
        }

        public Builder(CertManagerPrometheusServiceMonitorArgs defaults) {
            $ = new CertManagerPrometheusServiceMonitorArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder interval(@Nullable Output<String> interval) {
            $.interval = interval;
            return this;
        }

        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
        }

        public Builder prometheusInstance(@Nullable Output<String> prometheusInstance) {
            $.prometheusInstance = prometheusInstance;
            return this;
        }

        public Builder prometheusInstance(String prometheusInstance) {
            return prometheusInstance(Output.of(prometheusInstance));
        }

        public Builder string(@Nullable Output<String> string) {
            $.string = string;
            return this;
        }

        public Builder string(String string) {
            return string(Output.of(string));
        }

        public Builder targetPort(@Nullable Output<Integer> targetPort) {
            $.targetPort = targetPort;
            return this;
        }

        public Builder targetPort(Integer targetPort) {
            return targetPort(Output.of(targetPort));
        }

        public CertManagerPrometheusServiceMonitorArgs build() {
            return $;
        }
    }

}

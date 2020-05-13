# coding: utf-8

# # Consumer benchmarks

# ## Tasks
# * measure the speed of consuming messages from Kafka broker
# * measure speed of all steps performed during consume message operation
# * compute throughput - number of consumable messages per second (worst, best, average scenarios)
# * compute possible speedup achievable by using multiple consumers

# ## Preparation steps
# * 100000 messages were sent to Kafka broker into selected topic (in advance)
# * consumer has been updated to print durations into log files
# * storage has been set to be local (configurable PSQL or SQLite)

# ## Measurement steps
# * aggregator was started, all messages consumed, then stopped
# * log were redirected into text file
# * then log were transformed into two CSV files used below

# ## Machine used to run benchmarks
# ```
# 
# Architecture:        x86_64
# CPU op-mode(s):      32-bit, 64-bit
# Byte Order:          Little Endian
# CPU(s):              8
# On-line CPU(s) list: 0-7
# Thread(s) per core:  2
# Core(s) per socket:  4
# Socket(s):           1
# NUMA node(s):        1
# Vendor ID:           GenuineIntel
# CPU family:          6
# Model:               94
# Model name:          Intel(R) Core(TM) i7-6820HQ
#                      CPU @ 2.70GHz
# Stepping:            3
# CPU MHz:             900.222
# CPU max MHz:         3600.0000
# CPU min MHz:         800.0000
# BogoMIPS:            5424.00
# Virtualization:      VT-x
# L1d cache:           32K
# L1i cache:           32K
# L2 cache:            256K
# L3 cache:            8192K
# NUMA node0 CPU(s):   0-7
# ```

# ## Main results
# 
# Time was measured by the `time` tool on command line. Number of messages in
# Kafka topic was known in advance. So it is only needed to compute time in
# seconds (trivial) and average number of messages consumed per second:

number_of_consumed_messages=100000
time_in_minutes=26
time_in_seconds=time_in_minutes*60
messages_per_second=number_of_consumed_messages/time_in_seconds


# Average (rounded) number of messages consumed per second and per minute is:

print("Per second: ", int(messages_per_second))
print("Per minute: ", int(messages_per_second*60))


# ### Observations
# * one thread was used by aggregator (expected)
# * just 40% CPU utilization by aggregator process
# * rest (60%) spent by I/O operations
# * -> I/O (DB I/O + Kafka broker I/O basically are limiting factors)

# # Detailed behavior of consumer
# It is also possible to analyze log files (or rather CSV files generated from
# log files). We will use Pandas, Numpy, and Matplotlib libraries here

# ## Initialization part

# we are going to display graphs and work with data frames
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

# let's display all graphs without the need to call .show()
get_ipython().run_line_magic('matplotlib', 'inline')


# ## Loading all data files with raw metrics
# Two CSV files were prepared. `consumer_durations.csv` contains just whole duration and offset, nothing else:

# this CSV file contains just whole duration per message (ms) + message offset (int64 value)
durations=pd.read_csv("consumer_durations.csv")

# observe first ten items taken from this file
durations.head(10)


# Second file is named `consumer_steps_durations.csv`. It contains five values
# measured for each consumed message:

# 1. time to read message from Kafka topic
# 1. time to check if the message is correct
# 1. time to check if it is possible to marshall JSON stored in the message
# 1. time to check timestamp
# 1. time to store message body into DB storage

# this file is a bit more complicated - it contains duration of all 5 steps (in ns)
duration_steps=pd.read_csv("consumer_steps_durations.csv")


# first ten items taken from this file
duration_steps.head(10)


# ## Data statistic

# CSV files have been consumed and transformed into DataFrames, so it is
# possible to gather some statistic and display charts.

# let's compute average, best and worst durations (in ms) etc.
durations.describe()

# would be nice to display some graphs as well, especially for overall duration
durations["Duration"].plot()

# ## Detailed results for first 500 messages

# Please note that first x1000 messages are usually processed a bit faster
# compared to overall average! This is because garbage collector does not have
# to be started frequently during warmup and Go programs are not JITted.

# statistic (average, worst, best) for 5 steps for process each message
duration_steps.describe()

# again, plot the behaviour over time
duration_steps.plot()

# we can see that DB store is the most time demanding operation

# let's display relative times for each processing step
duration_steps.describe().transpose()["mean"].plot.pie(figsize=(6,6))

# ## Possible speedup - Amdahl's law

# It would be possible to perform first four steps in parallel. So let's
# compute if its worth it and which speedup is possible

# again, look at steps
duration_steps.describe()


# We can display stats/speedup for average, worst, and best scenarios. Average
# might be appropriate for the first version of this benchmark

# let's retrieve means for all five steps
means = duration_steps.describe().transpose()["mean"]

# the first four steps can be (in theory) made parallel
parallel_part = means["Read"]+means["Whitelisting"]+means["Marshalling"]+means["Time check"]
print("Parallel:", parallel_part, "ns")

# last step can be parallelized just in thery - in fact I/O is the bottleneck there
sequence_part = means["DB store"]
print("Sequence:", sequence_part, "ns")

# compute parameters for Amdahl's law
p=parallel_part/sequence_part
print("Ratio:", p)

# throughput for one pod/one CPU
t1 = 1000000/(parallel_part+sequence_part)
print("Throughput for 1 pod:", t1, "per second")

# now compute and display possible speedup for 2..32 CPUs/pods
s=np.arange(1, 33, 1)

# possible throughputs for 1..32 CPUs/pods
t=t1*1/(1-p+p/s)

# the best value for 32 CPUs/pods
print(t)

# display the graph
plt.rcParams["figure.figsize"] = (10,5)
fig=plt.figure()
plt.plot(s,t )
plt.show()

# looks like that even with 32 pods/CPUs (that is really large number of pods)
# we can process at most ~143 messages per second
per_second=143

# let's compute peak values per minute, per hour and per day
per_minute=per_second*60
per_hour=per_minute*60
per_day=per_hour*24
print("Per second", per_second)
print("Per minute", per_minute)
print("Per hour  ", per_hour)
print("Per day   ", per_day)

# ## Real expectations
# i.e. How much messages we have to process per given timeframe (day, hour,
# minute, second)?

# ### Loading all data files with raw metrics

# The following data file contains precise timestamps when input data were
# captured. We have to specify, that the first column needs to be parsed like a
# date (it can be done automatically, but sometimes it does not work
# correctly).
upload_timestamps=pd.read_csv("upload_timestamps_2020_04.csv", parse_dates=[0])

# Let's check the content of such data by displaying first ten records read
# from CSV file
upload_timestamps.head()



# ### Total uploads of insights raw data per day

# We can resample input data into one day buckets
by_day = upload_timestamps.resample('1D', on='Timestamp').count()
by_day.describe()

# display graph with measured total uploads per day
by_day_plot = by_day.plot(title="Total uploads per day",legend=None, kind="bar")



# ### Total uploads of insights raw data per hour

# The same operation can be done, but for 1 hour buckets
by_hour = upload_timestamps.resample('60min', on='Timestamp').count()
by_hour[:-1].describe()

# display graph with measured total uploads per hour
by_hour_plot = by_hour[:-1].plot(title="Total uploads per hour",legend=None)



# ### Total uploads of insights raw data per minute

# We can resample input data into 1 minute buckets
by_minute = upload_timestamps.resample('1min', on='Timestamp').count()
by_minute.describe()

# display graph with measured total uploads per hour
by_minute_plot = by_minute.plot(title="Total uploads per minute",legend=None)



# ### Total uploads of insights raw data per second

# It is possible to resample input data into 1 second buckets
by_second = upload_timestamps.resample('1s', on='Timestamp').count()
by_second.describe()

# display graph with measured total uploads per hour
by_second_plot = by_second.plot(title="Total uploads per second",legend=None)


# ## Conclusion

# Let's compare number of messages measured in production with the peak ratio
# (maximum number of messages that can be processed by using parallel pods)
per_second_stat=by_second.describe().transpose()
mean_value=per_second_stat["mean"].values[0]
worst_value=per_second_stat["max"].values[0]
best_value=per_second_stat["min"].values[0]

print("Average scenario: ", per_second, mean_value)
print("Best scenario:    ", per_second, best_value)
print("Worst scenario:   ", per_second, int(worst_value))


# # Aggregator memory consumption

# We also need to look how much memory is allocated by `aggregator` process.
# This process exposes metrics (as many other applications written in Go) that
# can be simply read with some frequency (ten seconds by default) and stored
# into CSV file named `memory_consumption.csv`. The following metrics are
# gathered and stored into CSV:

exported_metrics = (
        "go_gc_duration_seconds_sum",
        "go_gc_duration_seconds_count",
        "go_memstats_alloc_bytes",
        "go_memstats_sys_bytes",
        "go_memstats_mallocs_total",
        "go_memstats_frees_total",
        )

# Now it is possible to read file that contains memory consumption 
memory=pd.read_csv("memory_consumption.csv")

# Let's look at first 10 records just to see how values are stored
memory.head()

# And display graph with results
memory.plot(figsize=(10,30), grid=True, subplots=True)

# ## Conclusion

# Memory consumption is pretty low (8MB heap size) and - which is more
# important - it seems to be very stable over time. Also number of GC calls is
# low and does not cause slowdown of the whole process.

# finito

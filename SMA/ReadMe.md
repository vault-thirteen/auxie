# SMA

## Simple Moving Average

This is an example of simple moving average calculator using the built-in 
golang's list type.

This calculator uses an optimized calculation method which in its turn 
remembers the last calculated SMA value.

### Notes

1. In statistics, a **moving average** (**rolling average** or 
**running average**) is a calculation to analyze data points by creating a 
series of averages of different subsets of the full data set. It is also called 
a **moving mean** (**MM**) or **rolling mean** and is a type of finite impulse 
response filter. Variations include: simple, cumulative, or weighted forms.  
Source: https://en.wikipedia.org/wiki/Moving_average


2. In financial applications a **simple moving average** (**SMA**) is the 
unweighted mean of the previous `k` data-points.    
Source: https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average


3. In mathematics and statistics, the **arithmetic mean**, or simply the mean 
or the average (when the context is clear), is the sum of a collection of 
numbers divided by the count of numbers in the collection. The collection is 
often a set of results of an experiment or an observational study, or 
frequently a set of results from a survey. The term "arithmetic mean" is 
preferred in some contexts in mathematics and statistics, because it helps 
distinguish it from other means, such as the geometric mean and the harmonic 
mean.    
Source: https://en.wikipedia.org/wiki/Arithmetic_mean

### General Formula
![](https://wikimedia.org/api/rest_v1/media/math/render/svg/a608544726b8de1c3de562245ff0d1cd3d0efad6)

### Formula for Calculating Next SMA Value
This formula is used for calculating the next SMA value using the previous SMA value.
![](https://wikimedia.org/api/rest_v1/media/math/render/svg/cde134385756e2ead2222e4559f4a9128ab7eb52)

#!/usr/bin/env ruby
# frozen_string_literal: true
# Codeforces 926I — Alarm Clock (Ruby 3)

DAY_MINS = 24 * 60

def parse_time(s)
  h, m = s.split(':').map(&:to_i)
  h * 60 + m
end

def format_duration(mins)
  format('%02d:%02d', mins / 60, mins % 60)
end

def solve(alarms)
  return '23:59' if alarms.size == 1

  times = alarms.map { |s| parse_time(s) }.sort
  n = times.size

  best = (0...n).map do |i|
    j = (i + 1) % n
    diff = times[j] - times[i]
    diff += DAY_MINS if diff.negative?
    diff - 1
  end.max

  format_duration(best)
end

def drive(input)
  lines = input.each_line.map(&:strip).reject(&:empty?)
  n = lines[0].to_i
  alarms = lines[1, n]
  solve(alarms)
end

if $PROGRAM_NAME == __FILE__
  puts drive($stdin.read)
end

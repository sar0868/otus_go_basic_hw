package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcStatisticsChain(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name  string
		args  args
		level string
		want  Statistic
	}{
		{
			name: "statistic INFO module1=100% all=100%",
			args: args{[]string{
				"2025-01-09 14:59:48.7138 INFO module1 messages",
				"2025-01-09 14:59:48.7138 INFO module1 messages",
				"2025-01-09 14:59:48.7138 INFO module1 messages",
			}},
			level: "INFO",
			want: Statistic{
				level: "INFO",
				modules: map[string]int{
					"module1": 100,
				},
				all: 100,
			},
		},
		{
			name: "statistic INFO module1=50%, module2=25% all=33%",
			args: args{[]string{
				"2025-01-09 14:59:48.7138 INFO module1 messages",
				"2025-01-09 14:59:48.7138 INFO module2 messages",
				"2025-01-09 14:59:48.7138 WARN module1 messages",
				"2025-01-09 14:59:48.7138 WARN module2 messages",
				"2025-01-09 14:59:48.7138 FATAL module2 messages",
				"2025-01-09 14:59:48.7138 DEBUG module2 messages",
			}},
			level: "INFO",
			want: Statistic{
				level: "INFO",
				modules: map[string]int{
					"module1": 50,
					"module2": 25,
				},
				all: 33,
			},
		},
		{
			name: "statistic INFO module1=50%, module2=25% all=25%",
			args: args{[]string{
				"2025-01-09 14:59:48.7138 INFO module1 messages",
				"2025-01-09 14:59:48.7138 WARN module3 messages",
				"2025-01-09 14:59:48.7138 DEBUG module4 messages",
				"2025-01-09 14:59:48.7138 INFO module2 messages",
				"2025-01-09 14:59:48.7138 WARN module1 messages",
				"2025-01-09 14:59:48.7138 WARN module2 messages",
				"2025-01-09 14:59:48.7138 FATAL module2 messages",
				"2025-01-09 14:59:48.7138 DEBUG module2 messages",
			}},
			level: "INFO",
			want: Statistic{
				level: "INFO",
				modules: map[string]int{
					"module1": 50,
					"module2": 25,
				},
				all: 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan string)
			out := make(chan Statistic)
			go CalcStatistic(ch, tt.level, out)
			for _, el := range tt.args.data {
				ch <- el
			}
			close(ch)
			result := <-out
			close(out)
			assert.Equal(t, result, tt.want)
		})
	}
}

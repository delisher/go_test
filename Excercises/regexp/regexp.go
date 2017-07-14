package main

import (	
	"fmt"
	"regexp"
)

const (
	// Separator := "#{"\n"*3}#{"-"*100}#{"\n"*3}"
	// Mini_separator, _ := "\n"*3
)

func main() {
	tst := regexp.MustCompile("p([a-z]+)ch")
	rrr := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	ErrorsRE := regexp.MustCompile(`(\[ERROR\]|\[FATAL\]).+((\n.+){0,2})`)
	ErrorNameRE := regexp.MustCompile(`(?:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2},\d{3}\s).{10,30}`)
	DateRE := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}`)
	DrvRE := regexp.MustCompile(`\[execute_sql_script\]`)

	rexps := []*regexp.Regexp{tst, rrr, ErrorsRE, ErrorNameRE, DateRE, DrvRE}
	strngs := []string{
		"2017-06-26 00:00:17,086 [ERROR] -- Архивация.self_архивировать_TimedLinks: Вызов: nil.< [2017-06-12 00:00:00 +0300]  ",
		"2017-06-25 15:00:00,159 [ERROR] -- [OracleRawDataProvider#handle_error] Ошибка вставки объкта в таблицу Web_Eff_fifo_hist универсального хранилища;",
		"TypeError: for method PreparedStatement.setDouble expected [class int, class double]; got: [java.lang.Integer,java.math.BigDecimal]; error: null",
		"	file:/opt/ruby_projects/rvec/server/rvec_server/modules/ru-programpark-rvec-nb-server.jar!/db_data_provider.rb:1933:in `uobjectAccount_insert'",
		"	org/jruby/RubyHash.java:1371:in `each_pair'",
		"	file:/opt/ruby_projects/rvec/server/rvec_server/modules/ru-programpark-rvec-nb-server.jar!/db_data_provider.rb:1914:in `uobjectAccount_insert'",
		"	file:/opt/ruby_projects/rvec/server/rvec_server/modules/ru-programpark-rvec-nb-server.jar!/db_data_provider.rb:1909:in `uobjectAccount_insert'",
		"	org/jruby/RubyBasicObject.java:1721:in `__send__'",
		"	org/jruby/RubyKernel.java:2213:in `send'",
		"	file:/opt/ruby_projects/rvec/server/rvec_server/modules/ru-programpark-rvec-nb-server.jar!/raw_data_provider.rb:332:in `dispatch'",
		"	file:/opt/ruby_projects/rvec/server/rvec_server/modules/ru-programpark-rvec-nb-server.jar!/raw_data_provider.rb:224:in `start_save_thread'",
}
	for _, s := range strngs {
		fmt.Printf("%v =============================\n", s)
		for _, r := range rexps {
			fmt.Printf("\t`%v` =============================\n", r)
			fmt.Printf("\t\t`r.MatchString(s)` => %v\n", r.MatchString(s))
			fmt.Printf("\t\t`r.FindString(s)` => %v\n", r.FindString(s))
			fmt.Printf("\t\t`r.FindStringIndex(s)` => %v\n", r.FindStringIndex(s))
			fmt.Printf("\t\t`r.FindStringSubmatch(s)` => %v\n", r.FindStringSubmatch(s))
			fmt.Printf("\t\t`r.FindStringSubmatchIndex(s)` => %v\n", r.FindStringSubmatchIndex(s))
			fmt.Printf("\t\t`r.FindAllString(s, -1)` => %v\n", r.FindAllString(s, -1))
			fmt.Printf("\t\t`r.FindAllStringSubmatchIndex(s, -1)` => %v\n", r.FindAllStringSubmatchIndex(s, -1))
			fmt.Printf("\t\t`r.FindAllString(s, 2)` => %v\n", r.FindAllString(s, 2))
			fmt.Printf("\t\t`r.Match([]byte(s))` => %v\n", r.Match([]byte(s)))
			fmt.Println("=============================", "\n" ,"\n")
		}
	}
}
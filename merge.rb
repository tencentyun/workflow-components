out_file = File.open(ARGV[1], "w")

lines = File.new(ARGV[0]).each_line.map(&:chomp).each_slice(10) {|lines| 
  out_file.puts lines.join(' ')
}

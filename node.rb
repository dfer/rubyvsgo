require 'rubygems'
require 'digest/sha1'

# Стандартный вариант
def node0(text)
	return Integer('0x'+Digest::SHA1.hexdigest(text))%8
end

# Именно этот вариант сравниваем с node1 на Go.
def node1(text)
	s = Digest::SHA1.hexdigest(text)
	return Integer('0x'+s[s.size-1, 1]) % 8
end

# Этот вариант быстрее еще на 5% по сравнению с node1 на Ruby
def node11(text)
	s = Digest::SHA1.hexdigest(text)
	return ('0x'+s[s.size-1, 1]).hex % 8
end

# ---------- node v0 -----------------
date_from = Time.now.to_f
for i in 1..1000000 do 
	node0("user:"+i.to_s)
end
date_to = Time.now.to_f

delta = date_to - date_from

puts 'node0:'+delta.to_s

# ---------- node v1 -----------------
date_from = Time.now.to_f
for i in 1..1000000 do 
	node1("user:"+i.to_s)
end
date_to = Time.now.to_f

delta = date_to - date_from

puts 'node1:'+delta.to_s

# ---------- node v11 -----------------
date_from = Time.now.to_f
for i in 1..1000000 do 
	node11("user:"+i.to_s)
end
date_to = Time.now.to_f

delta = date_to - date_from

puts 'node1:'+delta.to_s
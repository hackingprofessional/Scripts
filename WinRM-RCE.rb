require 'winrm'

conn = WinRM::Connection.new( 
  endpoint: 'https://10.10.10.62:5555/wsman',
  transport: :ssl,
  user: 'WebUser',
  password: 'M4ng£m£ntPa55',
  :no_ssl_peer_verification => true
)

command=""

conn.shell(:powershell) do |shell|
    until command == "exit\n" do
        print "PS > "
        command = gets        
        output = shell.run(command) do |stdout, stderr|
            STDOUT.print stdout
            STDERR.print stderr
        end
    end    
    puts "Exiting with code #{output.exitcode}"
end


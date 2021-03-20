import time
from datetime import datetime, timedelta
import os
import sys

def primary(start_counter, counter_path, backup_path):
    print('[+] PRIMARY STARTED')
    counter = start_counter
    last_update_backup = datetime.now()

    while True:
        counter += 1
        print("PRIMARY: {} {}".format(counter, datetime.now()))
        
        f_counter = open(counter_path, 'w')
        f_counter.write("{}@{}".format(counter, datetime.now()))
        f_counter.close()
        
        # Check if the backup is alive. Create new if so
        f_backup = open(backup_path, 'r')
        line = f_backup.read()
        f_backup.close()
        
        try:
            timestamp = datetime.strptime(line, '%Y-%m-%d %H:%M:%S.%f')
        except ValueError:
            pass
        
        if last_update_backup != timestamp:
            last_update_backup = timestamp
        
        if datetime.now() - last_update_backup > timeout:
            print("[-] BACKUP TIMED OUT")
            create_backup(backup_path)
            
        time.sleep(1)
  
def create_backup(backup_path):
    # Write timestamp that backup is alive
    f_backup = open(backup_path, 'w')
    f_backup.write("{}".format(datetime.now()))
    f_backup.close()
    
    f = open('/Users/Jorgen/Google_Drive/NTNU/2021_var/sanntidsprogrammering/exercises/exercise-7-jorgenmyrvold/terminal_script.txt', 'r')
    script = f.read()
    f.close()
    os.system(script) 
    print('[+] BACKUP CREATED')          

def count(is_primary, counter_path, backup_path, timeout):
    last_update = datetime.now()
    counter = 0
    
    if is_primary:
        create_backup(backup_path)
        primary(0, counter_path, backup_path)
    
    while True:
        f_count = open(counter_path, 'r')
        line = f_count.read()
        f_count.close()
        
        try:
            args = line.split('@')
            counter = int(args[0])
            timestamp = datetime.strptime(args[1], '%Y-%m-%d %H:%M:%S.%f')
        except ValueError:
            pass
        
        if last_update != timestamp:
            last_update = timestamp
            
        if datetime.now() - last_update > timeout:
            print("[-] PRIMARY TIMED OUT")
            create_backup(backup_path)          
            break
        
        print('BACKUP:', counter)
        
        ## Write to file that backup is alive
        f_backup = open(backup_path, 'w')
        f_backup.write("{}".format(datetime.now()))
        f_backup.close()
        
        time.sleep(1)
        
    ## Take over as primary
    primary(counter, counter_path, backup_path)



if __name__ == "__main__":
    timeout = timedelta(seconds=3)
    counter_path = '/Users/Jorgen/Google_Drive/NTNU/2021_var/sanntidsprogrammering/exercises/exercise-7-jorgenmyrvold/count.txt'
    backup_path = '/Users/Jorgen/Google_Drive/NTNU/2021_var/sanntidsprogrammering/exercises/exercise-7-jorgenmyrvold/backup_alive.txt'
    is_primary = int(sys.argv[1])
    
    count(is_primary, counter_path, backup_path, timeout)
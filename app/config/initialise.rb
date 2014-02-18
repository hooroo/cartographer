app_path = File.expand_path(File.join('..', '..'), __FILE__)

APP_FILES = Dir[File.expand_path(File.join(app_path, '**', '*.rb'), __FILE__)].sort

APP_FILES.each do |file|
  require file
end

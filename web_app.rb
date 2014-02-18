require 'sinatra/base'
require 'sinatra/reloader'
require 'better_errors'
require 'json'

# require_relative './config/initialise'

class WebApp < Sinatra::Base

  configure :development do
    register Sinatra::Reloader

    # lib_path = File.expand_path(File.join('..', 'lib', 'transcoder_processor', '**', '*.rb'), __FILE__)

    # Dir[lib_path].each do |file|
    #   also_reload file
    # end

    use BetterErrors::Middleware
  end

  get '/' do
    erb :index
  end
end

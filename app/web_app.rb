require 'sinatra/base'
require 'sinatra/reloader'
require 'better_errors'
require 'json'

require_relative './config/initialise'
require_relative './presenters/node_presenter'
require_relative './presenters/report_presenter'

module Cartographer
  class WebApp < Sinatra::Base

    configure :development do
      register Sinatra::Reloader

      APP_FILES.each { |file| also_reload(file) }

      use BetterErrors::Middleware
    end

    get '/' do
      @nodes = get_nodes

      erb :index
    end

    private

      # TODO: Make this real!
      #
      def get_nodes
        nodes = JSON.parse('[
          {
            "system": { "node_name": "host.local", "ip_address": "192.168.0.1" },
            "reports": {
              "unicorn": { "workers": "4", "memory_used": "4096" },
              "nginx": { "workers": "8", "memory_used": "8192" }
            }
          }
        ]')

        nodes.map { |x| Presenters::NodePresenter.new(x) }
      end

  end
end

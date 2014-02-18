require 'hashie'

module Cartographer
  module Presenters
    class ReportPresenter

      attr_reader :identifier

      def initialize identifier, json
        @identifier = identifier
        @data = Hashie::Mash.new(json)
      end

      def detail
        data.map { |k, v| "#{k}=#{v}" }.join(', ')
      end

      def status
        # warning, danger
        'success'
      end

      private

        attr_reader :data

    end
  end
end

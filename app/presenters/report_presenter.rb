module Cartographer
  module Presenters
    class ReportPresenter

      def initialize json
        @json = json
      end

      private

        attr_reader :json

    end
  end
end

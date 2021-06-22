class ApiV1::AnnouncementsController < ApplicationController

  def index
    @announcements = Announcement.order(published_at: :desc)
  end
end

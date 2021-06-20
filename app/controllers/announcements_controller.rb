require 'net/http'

class AnnouncementsController < ApplicationController

  before_action :set_announcement, only: %i(edit update destroy)

  URI = URI.parse(ENV['PORTFOLIO_FRONT_DEPLOY_HOOKS_URL'])

  def index
    @announcements = Announcement.order(created_at: :desc)
  end

  def new
    @announcement = Announcement.new
  end

  def create
    Announcement.create!(allowed_params)
    if Rails.env.production?
      response = Net::HTTP.get_response(URI)
    end

    redirect_to announcements_path
  rescue => e
    redirect_back fallback_location: announcements_path
  end

  def edit; end

  def update
    @announcement.update!(allowed_params)
    if Rails.env.production?
      response = Net::HTTP.get_response(URI)
    end

    redirect_to announcements_path
  rescue => e
    redirect_back fallback_location: announcements_path
  end

  def destroy
    @announcement.destroy!
    if Rails.env.production?
      response = Net::HTTP.get_response(URI)
    end

    redirect_to announcements_path
  rescue => e
    redirect_back fallback_location: announcements_path
  end


  private

  def set_announcement
    @announcement = Announcement.find(params[:id])
  end

  def allowed_params
    params.require(:announcement).permit(
      :content,
      :link_url,
      :debug,
      :created_at,
      :updated_at
    )
  end
end
